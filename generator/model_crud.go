package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-courier/codegen"
	"github.com/go-courier/packagesx"
	"github.com/kunlun-qilian/sqlx/v3/builder"
)

func (m *Model) snippetSetDeletedAtIfNeedForFieldValues(file *codegen.File) codegen.Snippet {
	if m.HasDeletedAt {
		return codegen.Expr(`if _, ok := fieldValues["`+m.FieldKeyDeletedAt+`"]; !ok {
			fieldValues["`+m.FieldKeyDeletedAt+`"] = ?(`+file.Use("time", "Now")+`())
		}`,
			m.FieldType(file, m.FieldKeyDeletedAt),
		)
	}
	return nil
}

func (m *Model) snippetSetCreatedAtIfNeed(file *codegen.File) codegen.Snippet {
	if m.HasCreatedAt {
		return codegen.Expr(`
if m.`+m.FieldKeyCreatedAt+`.IsZero() {
	m.`+m.FieldKeyCreatedAt+` = ?(`+file.Use("time", "Now")+`())
}
`,
			m.FieldType(file, m.FieldKeyCreatedAt),
		)
	}

	return nil
}

func (m *Model) snippetSetUpdatedAtIfNeed(file *codegen.File) codegen.Snippet {
	if m.HasUpdatedAt {
		return codegen.Expr(`
if m.`+m.FieldKeyUpdatedAt+`.IsZero() {
	m.`+m.FieldKeyUpdatedAt+` = ?(`+file.Use("time", "Now")+`())
}
`,
			m.FieldType(file, m.FieldKeyUpdatedAt),
		)
	}

	return codegen.Expr("")
}

func (m *Model) snippetSetUpdatedAtIfNeedForFieldValues(file *codegen.File) codegen.Snippet {
	if m.HasUpdatedAt {
		return codegen.Expr(`
if _, ok := fieldValues[?]; !ok {
	fieldValues[?] = ?(?())
}
`,
			codegen.Val(m.FieldKeyUpdatedAt),
			codegen.Val(m.FieldKeyUpdatedAt),
			m.FieldType(file, m.FieldKeyUpdatedAt),
			codegen.Id(file.Use("time", "Now")))
	}
	return nil
}

func (m *Model) WriteCreate(file *codegen.File) {
	file.WriteBlock(
		codegen.Func(codegen.Var(
			codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
			Named("Create").
			MethodOf(codegen.Var(m.PtrType(), "m")).
			Return(codegen.Var(codegen.Error)).
			Do(
				m.snippetSetCreatedAtIfNeed(file),
				m.snippetSetUpdatedAtIfNeed(file),

				codegen.Expr(`
_, err := db.ExecExpr(?(db, m, nil))
return err
`,
					codegen.Id(file.Use("github.com/kunlun-qilian/sqlx/v3", "InsertToDB")),
				),
			),
	)

	if len(m.Keys.UniqueIndexes) > 0 {

		file.WriteBlock(
			codegen.Func(
				codegen.Var(codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db"),
				codegen.Var(codegen.Slice(codegen.String), "updateFields"),
			).
				Named("CreateOnDuplicateWithUpdateFields").
				MethodOf(codegen.Var(m.PtrType(), "m")).
				Return(codegen.Var(codegen.Error)).
				Do(
					codegen.Expr(`
if len(updateFields) == 0 {
	panic(`+file.Use("fmt", "Errorf")+`("must have update fields"))
}
`),

					m.snippetSetCreatedAtIfNeed(file),
					m.snippetSetUpdatedAtIfNeed(file),

					codegen.Expr(`
fieldValues := `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "FieldValuesFromStructByNonZero")+`(m, updateFields...)
`),
					func() codegen.Snippet {
						if m.HasAutoIncrement {
							return codegen.Expr(
								`delete(fieldValues, ?)`,
								file.Val(m.FieldKeyAutoIncrement),
							)
						}
						return nil
					}(),

					codegen.Expr(`
table := db.T(m)

cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

fields := make(map[string]bool, len(updateFields))
for _, field := range updateFields {
	fields[field] = true
}
`),
					codegen.Expr(`
for _, fieldNames := range m.UniqueIndexes() {
	for _, field := range fieldNames {
		delete(fields, field)
	}
}

if len(fields) == 0 {
	panic(`+file.Use("fmt", "Errorf")+`("no fields for updates"))
}

for field := range fieldValues {
	if !fields[field] {
		delete(fieldValues, field)
	}
}

additions := `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Additions")+`{}

switch db.Dialect().DriverName() {
case "mysql":	
	additions = append(additions, `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "OnDuplicateKeyUpdate")+`(table.AssignmentsByFieldValues(fieldValues)...))
case "postgres":
	indexes := m.UniqueIndexes()
	fields := make([]string, 0)
	for _, fs := range indexes {
		fields = append(fields, fs...)
	}
	indexFields, _ := db.T(m).Fields(fields...)

	additions = append(additions,
			`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "OnConflict")+`(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
}

additions = append(additions, `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`("User.CreateOnDuplicateWithUpdateFields"))

expr := `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Insert")+`().Into(table, additions...).Values(cols, vals...)

_, err := db.ExecExpr(expr)
return err
`,
						file.Val(m.StructName+".CreateOnDuplicateWithUpdateFields"),
						file.Val(m.StructName+".CreateOnDuplicateWithUpdateFields"),
					),
				),
		)
	}
}

func (m *Model) WriteDelete(file *codegen.File) {
	file.WriteBlock(
		codegen.Func(codegen.Var(
			codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
			Named("DeleteByStruct").
			MethodOf(codegen.Var(m.PtrType(), "m")).
			Return(codegen.Var(codegen.Error)).
			Do(
				codegen.Expr(`
_, err := db.ExecExpr(
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Delete")+`().
From(
	db.T(m),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Where")+`(m.ConditionByStruct(db)),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(?),
),
)

`, file.Val(m.StructName+".DeleteByStruct")),

				codegen.Return(codegen.Expr("err")),
			),
	)
}

func (m *Model) WriteByKey(file *codegen.File) {
	m.Table.Keys.Range(func(key *builder.Key, idx int) {
		fieldNames := key.Def.FieldNames

		fieldNamesWithoutEnabled := stringFilter(fieldNames, func(item string, i int) bool {
			if m.HasDeletedAt {
				return item != m.FieldKeyDeletedAt
			}
			return true
		})

		if m.HasDeletedAt && key.IsPrimary() {
			fieldNames = append(fieldNames, m.FieldKeyDeletedAt)
		}

		if key.IsUnique {
			{
				methodForFetch := createMethod("FetchBy%s", fieldNamesWithoutEnabled...)

				file.WriteBlock(
					codegen.Func(codegen.Var(
						codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
						Named(methodForFetch).
						MethodOf(codegen.Var(m.PtrType(), "m")).
						Return(codegen.Var(codegen.Error)).
						Do(
							codegen.Expr(`
table := db.T(m)

err := db.QueryExprAndScan(
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Select")+`(nil).
From(
	db.T(m),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Where")+`(`+toExactlyConditionFrom(file, fieldNames...)+`),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(?),
),
m,
)
`,
								file.Val(m.StructName+"."+methodForFetch),
							),

							codegen.Return(codegen.Expr("err")),
						),
				)

				methodForUpdateWithMap := createMethod("UpdateBy%sWithMap", fieldNamesWithoutEnabled...)

				file.WriteBlock(
					codegen.Func(
						codegen.Var(codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db"),
						codegen.Var(codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "FieldValues")), "fieldValues"),
					).
						Named(methodForUpdateWithMap).
						MethodOf(codegen.Var(m.PtrType(), "m")).
						Return(codegen.Var(codegen.Error)).
						Do(
							m.snippetSetUpdatedAtIfNeedForFieldValues(file),
							codegen.Expr(`
table := db.T(m)

result, err := db.ExecExpr(
	`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Update")+`(db.T(m)).
		Where(
			`+toExactlyConditionFrom(file, fieldNames...)+`,
			`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(?),
		).
		Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

if err != nil {
	return err
}

rowsAffected, _ := result.RowsAffected()
if rowsAffected == 0 {
  return m.`+methodForFetch+`(db)
}

return nil
`,
								file.Val(m.StructName+"."+methodForUpdateWithMap),
							),
						),
				)

				methodForUpdateWithStruct := createMethod("UpdateBy%sWithStruct", fieldNamesWithoutEnabled...)

				file.WriteBlock(
					codegen.Func(
						codegen.Var(codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db"),
						codegen.Var(codegen.Ellipsis(codegen.String), "zeroFields"),
					).
						Named(methodForUpdateWithStruct).
						MethodOf(codegen.Var(m.PtrType(), "m")).
						Return(codegen.Var(codegen.Error)).
						Do(
							codegen.Expr(`
fieldValues := ` + file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "FieldValuesFromStructByNonZero") + `(m, zeroFields...)
return m.` + methodForUpdateWithMap + `(db, fieldValues)
`),
						),
				)
			}

			{

				method := createMethod("FetchBy%sForUpdate", fieldNamesWithoutEnabled...)

				file.WriteBlock(
					codegen.Func(codegen.Var(
						codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
						Named(method).
						MethodOf(codegen.Var(m.PtrType(), "m")).
						Return(codegen.Var(codegen.Error)).
						Do(
							codegen.Expr(`
table := db.T(m)

err := db.QueryExprAndScan(
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Select")+`(nil).
From(
	db.T(m),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Where")+`(`+toExactlyConditionFrom(file, fieldNames...)+`),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "ForUpdate")+`(),
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(?),
),
m,
)
`,
								file.Val(m.StructName+"."+method),
							),

							codegen.Return(codegen.Expr("err")),
						),
				)
			}

			{
				methodForDelete := createMethod("DeleteBy%s", fieldNamesWithoutEnabled...)

				file.WriteBlock(
					codegen.Func(codegen.Var(
						codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
						Named(methodForDelete).
						MethodOf(codegen.Var(m.PtrType(), "m")).
						Return(codegen.Var(codegen.Error)).
						Do(
							codegen.Expr(`
table := db.T(m)

_, err := db.ExecExpr(
`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Delete")+`().
	From(db.T(m),
	`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Where")+`(`+toExactlyConditionFrom(file, fieldNames...)+`),
	`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(`+string(file.Val(m.StructName+"."+methodForDelete).Bytes())+`),
))
`,
								file.Val(m.StructName+"."+methodForDelete),
							),

							codegen.Return(codegen.Expr("err")),
						),
				)

				if m.HasDeletedAt {

					methodForSoftDelete := createMethod("SoftDeleteBy%s", fieldNamesWithoutEnabled...)

					file.WriteBlock(
						codegen.Func(codegen.Var(
							codegen.Type(file.Use("github.com/kunlun-qilian/sqlx/v3", "DBExecutor")), "db")).
							Named(methodForSoftDelete).
							MethodOf(codegen.Var(m.PtrType(), "m")).
							Return(codegen.Var(codegen.Error)).
							Do(
								codegen.Expr(`
table := db.T(m)

fieldValues := `+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "FieldValues")+`{}`),

								m.snippetSetDeletedAtIfNeedForFieldValues(file),
								m.snippetSetUpdatedAtIfNeedForFieldValues(file),

								codegen.Expr(`
_, err := db.ExecExpr(
	`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Update")+`(db.T(m)).
		Where(
			`+toExactlyConditionFrom(file, fieldNames...)+`,
			`+file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "Comment")+`(`+string(file.Val(m.StructName+"."+methodForSoftDelete).Bytes())+`),
		).
		Set(table.AssignmentsByFieldValues(fieldValues)...),
)

return err
`),
							),
					)
				}
			}
		}
	})
}

func (m *Model) WriteCRUD(file *codegen.File) {
	m.WriteCreate(file)
	m.WriteDelete(file)
	m.WriteByKey(file)
}

func toExactlyConditionFrom(file *codegen.File, fieldNames ...string) string {
	buf := &bytes.Buffer{}
	buf.WriteString(file.Use("github.com/kunlun-qilian/sqlx/v3/builder", "And"))
	buf.WriteString(`(
`)

	for _, fieldName := range fieldNames {
		buf.WriteString(fmt.Sprintf(`table.F("%s").Eq(m.%s),
		`, fieldName, fieldName))
	}

	buf.WriteString(`
)`)

	return buf.String()
}

func createMethod(method string, fieldNames ...string) string {
	return fmt.Sprintf(method, strings.Join(fieldNames, "And"))
}

func (m *Model) FieldType(file *codegen.File, fieldName string) codegen.SnippetType {
	if field, ok := m.Fields[fieldName]; ok {
		typ := field.Type().String()
		if strings.Contains(typ, ".") {
			importPath, name := packagesx.GetPkgImportPathAndExpose(typ)
			if importPath != m.TypeName.Pkg().Path() {
				return codegen.Type(file.Use(importPath, name))
			}
			return codegen.Type(name)
		}
		return codegen.BuiltInType(typ)
	}
	return nil
}
