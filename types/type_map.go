package types

import (
	"fmt"
	"go/ast"
	"io"
)

const Map string = "map"

func NewMap(key, value TypeDef) *typeMap {
	return &typeMap{key: key, value: value}
}

type typeMap struct {
	min   *string
	max   *string
	key   TypeDef
	value TypeDef
}

func (t typeMap) Type() string {
	return Map
}

func (t *typeMap) SetValidateTag(tag ValidatableTag) error {
	switch tag.Key() {
	case MapMinItemsKey:
		st := tag.(SimpleTag)
		t.min = &st.Param
	case MapMaxItemsKey:
		st := tag.(SimpleTag)
		t.max = &st.Param
	case MapKeyKey:
		scope := tag.(ScopeTag)
		for _, it := range scope.InnerTags {
			if err := t.key.SetValidateTag(it); err != nil {
				return fmt.Errorf("set item tags for key failed, tag %+v, err %s", it, err)
			}
		}
	case MapValueKey:
		scope := tag.(ScopeTag)
		for _, it := range scope.InnerTags {
			if err := t.value.SetValidateTag(it); err != nil {
				return fmt.Errorf("set item tags for value failed, tag %+v, err %s", it, err)
			}
		}
	default:
		return ErrUnusedTag
	}
	return nil
}

func (t typeMap) NeedGenerate() bool {
	return t.key.NeedGenerate() || t.value.NeedGenerate() || validMaxMin(t.max, t.min)
}

func (t typeMap) Generate(w io.Writer, cfg GenConfig, name Name) {
	if t.min != nil {
		if *t.min != "0" {
			fmt.Fprintf(w, "if len(%s) < %s {\n", name.Full(), *t.min)
			fmt.Fprintf(w, "    errs.AddFieldf(%s, \"less items than %s\")\n", name.LabelName(), *t.min)
			fmt.Fprintf(w, "}\n")
		}
	}
	if t.max != nil {
		fmt.Fprintf(w, "if len(%s) > %s {\n", name.Full(), *t.max)
		fmt.Fprintf(w, "    errs.AddFieldf(%s, \"more items than %s\")\n", name.LabelName(), *t.max)
		fmt.Fprintf(w, "}\n")
	}

	if t.key.NeedGenerate() || t.value.NeedGenerate() {
		kName := "k" + name.fieldName
		vName := "v" + name.fieldName
		fmt.Fprintf(w, "for %s, %s := range %s {\n", kName, vName, name.Full())
		fmt.Fprintf(w, "	_ = %s \n", kName)
		fmt.Fprintf(w, "	_ = %s \n", vName)
		cfg.AddImport("fmt")
		t.key.Generate(w, cfg, NewIndexedName(name.labelName, kName, kName, name.tagName))
		t.value.Generate(w, cfg, NewIndexedName(name.labelName, vName, vName, name.tagName))
		fmt.Fprintf(w, "}\n")
	}
}

func (t typeMap) Validate() error {
	if err := validateMinMax(
		t.min,
		t.max,
		func(min float64) error {
			if min < 0 {
				return fmt.Errorf("min map items can't be less than 0: %f", min)
			}
			return nil
		},
		func(max float64) error {
			if max < 0 {
				return fmt.Errorf("max map items can't be less than 0: %f", max)
			}
			return nil
		},
	); err != nil {
		return err
	}
	if err := t.key.Validate(); err != nil {
		return err
	}
	return t.value.Validate()
}

func genName(baseName string, key bool) string {
	postfix := "Value"
	if key {
		postfix = "Key"
	}
	return fmt.Sprintf("%s%s", baseName, postfix)
}

func (t typeMap) Expr() ast.Expr {
	return nil
}
