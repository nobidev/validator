package validator

import (
	"context"
	ut "github.com/go-playground/universal-translator"
)

var Default = New()

func SetTagName(name string) {
	Default.SetTagName(name)
}

func ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return Default.ValidateMapCtx(ctx, data, rules)
}

func ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return Default.ValidateMap(data, rules)
}

func RegisterTagNameFunc(fn TagNameFunc) {
	Default.RegisterTagNameFunc(fn)
}

func RegisterValidation(tag string, fn Func, callValidationEvenIfNull ...bool) error {
	return Default.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func RegisterValidationCtx(tag string, fn FuncCtx, callValidationEvenIfNull ...bool) error {
	return Default.RegisterValidationCtx(tag, fn, callValidationEvenIfNull...)
}

func registerValidation(tag string, fn FuncCtx, bakedIn bool, nilCheckable bool) error {
	return Default.registerValidation(tag, fn, bakedIn, nilCheckable)
}

func RegisterAlias(alias, tags string) {
	Default.RegisterAlias(alias, tags)
}

func RegisterStructValidation(fn StructLevelFunc, types ...interface{}) {
	Default.RegisterStructValidation(fn, types)
}

func RegisterStructValidationCtx(fn StructLevelFuncCtx, types ...interface{}) {
	Default.RegisterStructValidationCtx(fn, types)
}

func RegisterStructValidationMapRules(rules map[string]string, types ...interface{}) {
	Default.RegisterStructValidationMapRules(rules, types...)
}

func RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{}) {
	Default.RegisterCustomTypeFunc(fn, types)
}

func RegisterTranslation(tag string, trans ut.Translator, registerFn RegisterTranslationsFunc, translationFn TranslationFunc) error {
	return Default.RegisterTranslation(tag, trans, registerFn, translationFn)
}

func Struct(s interface{}) error {
	return Default.Struct(s)
}

func StructCtx(ctx context.Context, s interface{}) error {
	return Default.StructCtx(ctx, s)
}

func StructFiltered(s interface{}, fn FilterFunc) error {
	return Default.StructFiltered(s, fn)
}

func StructFilteredCtx(ctx context.Context, s interface{}, fn FilterFunc) error {
	return Default.StructFilteredCtx(ctx, s, fn)
}

func StructPartial(s interface{}, fields ...string) error {
	return Default.StructPartial(s, fields...)
}

func StructPartialCtx(ctx context.Context, s interface{}, fields ...string) error {
	return Default.StructPartialCtx(ctx, s, fields...)
}

func StructExcept(s interface{}, fields ...string) error {
	return Default.StructExcept(s, fields...)
}

func StructExceptCtx(ctx context.Context, s interface{}, fields ...string) error {
	return Default.StructExceptCtx(ctx, s, fields...)
}

func Var(field interface{}, tag string) error {
	return Default.Var(field, tag)
}

func VarCtx(ctx context.Context, field interface{}, tag string) error {
	return Default.VarCtx(ctx, field, tag)
}

func VarWithValue(field interface{}, other interface{}, tag string) error {
	return Default.VarWithValue(field, other, tag)
}

func VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) error {
	return Default.VarWithValueCtx(ctx, field, other, tag)
}
