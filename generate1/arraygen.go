package gen1

//go:generate bash -c "sed -e 's/XTYPEX/int/g' _array_template.go > array_int_gen.go"
//go:generate bash -c "sed -e 's/XTYPEX/string/g' _array_template.go > array_string_gen.go"
