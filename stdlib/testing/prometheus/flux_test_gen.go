// DO NOT EDIT: This file is autogenerated via the builtin command.

package prometheus

import (
	ast "github.com/influxdata/flux/ast"
	parser "github.com/influxdata/flux/internal/parser"
)

var FluxTestPackages = []*ast.Package{&ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 106,
					Line:   54,
				},
				File:   "histogramQuantile_test.flux",
				Source: "package prometheus_test\n\nimport \"experimental/prometheus\" \nimport \"testing\"\n\noption now = () => (2030-01-01T00:00:00Z)\n\ninData = \"\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,0,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,100,1\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,1,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,150,2\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,2,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,200,25\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,3,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,250,27\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,4,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,300,27\n\"\n\noutData = \"\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,string,double\n#group,false,false,true,true,true,true,true,false\n#default,_result,,,,,,,\n,result,table,_start,_stop,_measurement,_field,url,_value\n,,0,2018-05-22T13:00:00Z,2030-01-01T00:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,175\n\"\n\nt_histogramQuantile = (table=<-) =>\n    (table\n        |> range(start: 2018-05-22T13:00:00Z))\n        |> prometheus.histogramQuantile(quantile: 0.5)\n\ntest _histogramQuantile = () => \n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile})",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.OptionStatement{
			Assignment: &ast.VariableAssignment{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 42,
							Line:   6,
						},
						File:   "histogramQuantile_test.flux",
						Source: "now = () => (2030-01-01T00:00:00Z)",
						Start: ast.Position{
							Column: 8,
							Line:   6,
						},
					},
				},
				ID: &ast.Identifier{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 11,
								Line:   6,
							},
							File:   "histogramQuantile_test.flux",
							Source: "now",
							Start: ast.Position{
								Column: 8,
								Line:   6,
							},
						},
					},
					Name: "now",
				},
				Init: &ast.FunctionExpression{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 42,
								Line:   6,
							},
							File:   "histogramQuantile_test.flux",
							Source: "() => (2030-01-01T00:00:00Z)",
							Start: ast.Position{
								Column: 14,
								Line:   6,
							},
						},
					},
					Body: &ast.ParenExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 42,
									Line:   6,
								},
								File:   "histogramQuantile_test.flux",
								Source: "(2030-01-01T00:00:00Z)",
								Start: ast.Position{
									Column: 20,
									Line:   6,
								},
							},
						},
						Expression: &ast.DateTimeLiteral{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 41,
										Line:   6,
									},
									File:   "histogramQuantile_test.flux",
									Source: "2030-01-01T00:00:00Z",
									Start: ast.Position{
										Column: 21,
										Line:   6,
									},
								},
							},
							Value: parser.MustParseTime("2030-01-01T00:00:00Z"),
						},
					},
					Params: []*ast.Property{},
				},
			},
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 42,
						Line:   6,
					},
					File:   "histogramQuantile_test.flux",
					Source: "option now = () => (2030-01-01T00:00:00Z)",
					Start: ast.Position{
						Column: 1,
						Line:   6,
					},
				},
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   38,
					},
					File:   "histogramQuantile_test.flux",
					Source: "inData = \"\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,0,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,100,1\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,1,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,150,2\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,2,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,200,25\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,3,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,250,27\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,4,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,300,27\n\"",
					Start: ast.Position{
						Column: 1,
						Line:   8,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 7,
							Line:   8,
						},
						File:   "histogramQuantile_test.flux",
						Source: "inData",
						Start: ast.Position{
							Column: 1,
							Line:   8,
						},
					},
				},
				Name: "inData",
			},
			Init: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   38,
						},
						File:   "histogramQuantile_test.flux",
						Source: "\"\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,0,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,100,1\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,1,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,150,2\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,2,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,200,25\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,3,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,250,27\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,4,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,300,27\n\"",
						Start: ast.Position{
							Column: 10,
							Line:   8,
						},
					},
				},
				Value: "\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,0,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,100,1\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,1,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,150,2\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,2,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,200,25\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,3,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,250,27\n\n#datatype,string,long,dateTime:RFC3339,string,string,string,string,double\n#group,false,false,false,true,true,false,true,false\n#default,_result,,,,,,,\n,result,table,_time,_measurement,_field,url,le,_value\n,,4,2018-05-22T13:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,300,27\n",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 2,
						Line:   46,
					},
					File:   "histogramQuantile_test.flux",
					Source: "outData = \"\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,string,double\n#group,false,false,true,true,true,true,true,false\n#default,_result,,,,,,,\n,result,table,_start,_stop,_measurement,_field,url,_value\n,,0,2018-05-22T13:00:00Z,2030-01-01T00:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,175\n\"",
					Start: ast.Position{
						Column: 1,
						Line:   40,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 8,
							Line:   40,
						},
						File:   "histogramQuantile_test.flux",
						Source: "outData",
						Start: ast.Position{
							Column: 1,
							Line:   40,
						},
					},
				},
				Name: "outData",
			},
			Init: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 2,
							Line:   46,
						},
						File:   "histogramQuantile_test.flux",
						Source: "\"\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,string,double\n#group,false,false,true,true,true,true,true,false\n#default,_result,,,,,,,\n,result,table,_start,_stop,_measurement,_field,url,_value\n,,0,2018-05-22T13:00:00Z,2030-01-01T00:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,175\n\"",
						Start: ast.Position{
							Column: 11,
							Line:   40,
						},
					},
				},
				Value: "\n#datatype,string,long,dateTime:RFC3339,dateTime:RFC3339,string,string,string,double\n#group,false,false,true,true,true,true,true,false\n#default,_result,,,,,,,\n,result,table,_start,_stop,_measurement,_field,url,_value\n,,0,2018-05-22T13:00:00Z,2030-01-01T00:00:00Z,prometheus,prometheus_test_metric,http://prometheus.test,175\n",
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 55,
						Line:   51,
					},
					File:   "histogramQuantile_test.flux",
					Source: "t_histogramQuantile = (table=<-) =>\n    (table\n        |> range(start: 2018-05-22T13:00:00Z))\n        |> prometheus.histogramQuantile(quantile: 0.5)",
					Start: ast.Position{
						Column: 1,
						Line:   48,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 20,
							Line:   48,
						},
						File:   "histogramQuantile_test.flux",
						Source: "t_histogramQuantile",
						Start: ast.Position{
							Column: 1,
							Line:   48,
						},
					},
				},
				Name: "t_histogramQuantile",
			},
			Init: &ast.FunctionExpression{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 55,
							Line:   51,
						},
						File:   "histogramQuantile_test.flux",
						Source: "(table=<-) =>\n    (table\n        |> range(start: 2018-05-22T13:00:00Z))\n        |> prometheus.histogramQuantile(quantile: 0.5)",
						Start: ast.Position{
							Column: 23,
							Line:   48,
						},
					},
				},
				Body: &ast.PipeExpression{
					Argument: &ast.ParenExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 47,
									Line:   50,
								},
								File:   "histogramQuantile_test.flux",
								Source: "(table\n        |> range(start: 2018-05-22T13:00:00Z))",
								Start: ast.Position{
									Column: 5,
									Line:   49,
								},
							},
						},
						Expression: &ast.PipeExpression{
							Argument: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 11,
											Line:   49,
										},
										File:   "histogramQuantile_test.flux",
										Source: "table",
										Start: ast.Position{
											Column: 6,
											Line:   49,
										},
									},
								},
								Name: "table",
							},
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 46,
										Line:   50,
									},
									File:   "histogramQuantile_test.flux",
									Source: "table\n        |> range(start: 2018-05-22T13:00:00Z)",
									Start: ast.Position{
										Column: 6,
										Line:   49,
									},
								},
							},
							Call: &ast.CallExpression{
								Arguments: []ast.Expression{&ast.ObjectExpression{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 45,
												Line:   50,
											},
											File:   "histogramQuantile_test.flux",
											Source: "start: 2018-05-22T13:00:00Z",
											Start: ast.Position{
												Column: 18,
												Line:   50,
											},
										},
									},
									Properties: []*ast.Property{&ast.Property{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 45,
													Line:   50,
												},
												File:   "histogramQuantile_test.flux",
												Source: "start: 2018-05-22T13:00:00Z",
												Start: ast.Position{
													Column: 18,
													Line:   50,
												},
											},
										},
										Key: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 23,
														Line:   50,
													},
													File:   "histogramQuantile_test.flux",
													Source: "start",
													Start: ast.Position{
														Column: 18,
														Line:   50,
													},
												},
											},
											Name: "start",
										},
										Value: &ast.DateTimeLiteral{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 45,
														Line:   50,
													},
													File:   "histogramQuantile_test.flux",
													Source: "2018-05-22T13:00:00Z",
													Start: ast.Position{
														Column: 25,
														Line:   50,
													},
												},
											},
											Value: parser.MustParseTime("2018-05-22T13:00:00Z"),
										},
									}},
									With: nil,
								}},
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 46,
											Line:   50,
										},
										File:   "histogramQuantile_test.flux",
										Source: "range(start: 2018-05-22T13:00:00Z)",
										Start: ast.Position{
											Column: 12,
											Line:   50,
										},
									},
								},
								Callee: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 17,
												Line:   50,
											},
											File:   "histogramQuantile_test.flux",
											Source: "range",
											Start: ast.Position{
												Column: 12,
												Line:   50,
											},
										},
									},
									Name: "range",
								},
							},
						},
					},
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 55,
								Line:   51,
							},
							File:   "histogramQuantile_test.flux",
							Source: "(table\n        |> range(start: 2018-05-22T13:00:00Z))\n        |> prometheus.histogramQuantile(quantile: 0.5)",
							Start: ast.Position{
								Column: 5,
								Line:   49,
							},
						},
					},
					Call: &ast.CallExpression{
						Arguments: []ast.Expression{&ast.ObjectExpression{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 54,
										Line:   51,
									},
									File:   "histogramQuantile_test.flux",
									Source: "quantile: 0.5",
									Start: ast.Position{
										Column: 41,
										Line:   51,
									},
								},
							},
							Properties: []*ast.Property{&ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 54,
											Line:   51,
										},
										File:   "histogramQuantile_test.flux",
										Source: "quantile: 0.5",
										Start: ast.Position{
											Column: 41,
											Line:   51,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 49,
												Line:   51,
											},
											File:   "histogramQuantile_test.flux",
											Source: "quantile",
											Start: ast.Position{
												Column: 41,
												Line:   51,
											},
										},
									},
									Name: "quantile",
								},
								Value: &ast.FloatLiteral{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 54,
												Line:   51,
											},
											File:   "histogramQuantile_test.flux",
											Source: "0.5",
											Start: ast.Position{
												Column: 51,
												Line:   51,
											},
										},
									},
									Value: 0.5,
								},
							}},
							With: nil,
						}},
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 55,
									Line:   51,
								},
								File:   "histogramQuantile_test.flux",
								Source: "prometheus.histogramQuantile(quantile: 0.5)",
								Start: ast.Position{
									Column: 12,
									Line:   51,
								},
							},
						},
						Callee: &ast.MemberExpression{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 40,
										Line:   51,
									},
									File:   "histogramQuantile_test.flux",
									Source: "prometheus.histogramQuantile",
									Start: ast.Position{
										Column: 12,
										Line:   51,
									},
								},
							},
							Object: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 22,
											Line:   51,
										},
										File:   "histogramQuantile_test.flux",
										Source: "prometheus",
										Start: ast.Position{
											Column: 12,
											Line:   51,
										},
									},
								},
								Name: "prometheus",
							},
							Property: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 40,
											Line:   51,
										},
										File:   "histogramQuantile_test.flux",
										Source: "histogramQuantile",
										Start: ast.Position{
											Column: 23,
											Line:   51,
										},
									},
								},
								Name: "histogramQuantile",
							},
						},
					},
				},
				Params: []*ast.Property{&ast.Property{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 32,
								Line:   48,
							},
							File:   "histogramQuantile_test.flux",
							Source: "table=<-",
							Start: ast.Position{
								Column: 24,
								Line:   48,
							},
						},
					},
					Key: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 29,
									Line:   48,
								},
								File:   "histogramQuantile_test.flux",
								Source: "table",
								Start: ast.Position{
									Column: 24,
									Line:   48,
								},
							},
						},
						Name: "table",
					},
					Value: &ast.PipeLiteral{BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 32,
								Line:   48,
							},
							File:   "histogramQuantile_test.flux",
							Source: "<-",
							Start: ast.Position{
								Column: 30,
								Line:   48,
							},
						},
					}},
				}},
			},
		}, &ast.TestStatement{
			Assignment: &ast.VariableAssignment{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 106,
							Line:   54,
						},
						File:   "histogramQuantile_test.flux",
						Source: "_histogramQuantile = () => \n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile})",
						Start: ast.Position{
							Column: 6,
							Line:   53,
						},
					},
				},
				ID: &ast.Identifier{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 24,
								Line:   53,
							},
							File:   "histogramQuantile_test.flux",
							Source: "_histogramQuantile",
							Start: ast.Position{
								Column: 6,
								Line:   53,
							},
						},
					},
					Name: "_histogramQuantile",
				},
				Init: &ast.FunctionExpression{
					BaseNode: ast.BaseNode{
						Errors: nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 106,
								Line:   54,
							},
							File:   "histogramQuantile_test.flux",
							Source: "() => \n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile})",
							Start: ast.Position{
								Column: 27,
								Line:   53,
							},
						},
					},
					Body: &ast.ParenExpression{
						BaseNode: ast.BaseNode{
							Errors: nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 106,
									Line:   54,
								},
								File:   "histogramQuantile_test.flux",
								Source: "({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile})",
								Start: ast.Position{
									Column: 1,
									Line:   54,
								},
							},
						},
						Expression: &ast.ObjectExpression{
							BaseNode: ast.BaseNode{
								Errors: nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 105,
										Line:   54,
									},
									File:   "histogramQuantile_test.flux",
									Source: "{input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile}",
									Start: ast.Position{
										Column: 2,
										Line:   54,
									},
								},
							},
							Properties: []*ast.Property{&ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 42,
											Line:   54,
										},
										File:   "histogramQuantile_test.flux",
										Source: "input: testing.loadStorage(csv: inData)",
										Start: ast.Position{
											Column: 3,
											Line:   54,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 8,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "input",
											Start: ast.Position{
												Column: 3,
												Line:   54,
											},
										},
									},
									Name: "input",
								},
								Value: &ast.CallExpression{
									Arguments: []ast.Expression{&ast.ObjectExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 41,
													Line:   54,
												},
												File:   "histogramQuantile_test.flux",
												Source: "csv: inData",
												Start: ast.Position{
													Column: 30,
													Line:   54,
												},
											},
										},
										Properties: []*ast.Property{&ast.Property{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 41,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "csv: inData",
													Start: ast.Position{
														Column: 30,
														Line:   54,
													},
												},
											},
											Key: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 33,
															Line:   54,
														},
														File:   "histogramQuantile_test.flux",
														Source: "csv",
														Start: ast.Position{
															Column: 30,
															Line:   54,
														},
													},
												},
												Name: "csv",
											},
											Value: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 41,
															Line:   54,
														},
														File:   "histogramQuantile_test.flux",
														Source: "inData",
														Start: ast.Position{
															Column: 35,
															Line:   54,
														},
													},
												},
												Name: "inData",
											},
										}},
										With: nil,
									}},
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 42,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "testing.loadStorage(csv: inData)",
											Start: ast.Position{
												Column: 10,
												Line:   54,
											},
										},
									},
									Callee: &ast.MemberExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 29,
													Line:   54,
												},
												File:   "histogramQuantile_test.flux",
												Source: "testing.loadStorage",
												Start: ast.Position{
													Column: 10,
													Line:   54,
												},
											},
										},
										Object: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 17,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "testing",
													Start: ast.Position{
														Column: 10,
														Line:   54,
													},
												},
											},
											Name: "testing",
										},
										Property: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 29,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "loadStorage",
													Start: ast.Position{
														Column: 18,
														Line:   54,
													},
												},
											},
											Name: "loadStorage",
										},
									},
								},
							}, &ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 79,
											Line:   54,
										},
										File:   "histogramQuantile_test.flux",
										Source: "want: testing.loadMem(csv: outData)",
										Start: ast.Position{
											Column: 44,
											Line:   54,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 48,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "want",
											Start: ast.Position{
												Column: 44,
												Line:   54,
											},
										},
									},
									Name: "want",
								},
								Value: &ast.CallExpression{
									Arguments: []ast.Expression{&ast.ObjectExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 78,
													Line:   54,
												},
												File:   "histogramQuantile_test.flux",
												Source: "csv: outData",
												Start: ast.Position{
													Column: 66,
													Line:   54,
												},
											},
										},
										Properties: []*ast.Property{&ast.Property{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 78,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "csv: outData",
													Start: ast.Position{
														Column: 66,
														Line:   54,
													},
												},
											},
											Key: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 69,
															Line:   54,
														},
														File:   "histogramQuantile_test.flux",
														Source: "csv",
														Start: ast.Position{
															Column: 66,
															Line:   54,
														},
													},
												},
												Name: "csv",
											},
											Value: &ast.Identifier{
												BaseNode: ast.BaseNode{
													Errors: nil,
													Loc: &ast.SourceLocation{
														End: ast.Position{
															Column: 78,
															Line:   54,
														},
														File:   "histogramQuantile_test.flux",
														Source: "outData",
														Start: ast.Position{
															Column: 71,
															Line:   54,
														},
													},
												},
												Name: "outData",
											},
										}},
										With: nil,
									}},
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 79,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "testing.loadMem(csv: outData)",
											Start: ast.Position{
												Column: 50,
												Line:   54,
											},
										},
									},
									Callee: &ast.MemberExpression{
										BaseNode: ast.BaseNode{
											Errors: nil,
											Loc: &ast.SourceLocation{
												End: ast.Position{
													Column: 65,
													Line:   54,
												},
												File:   "histogramQuantile_test.flux",
												Source: "testing.loadMem",
												Start: ast.Position{
													Column: 50,
													Line:   54,
												},
											},
										},
										Object: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 57,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "testing",
													Start: ast.Position{
														Column: 50,
														Line:   54,
													},
												},
											},
											Name: "testing",
										},
										Property: &ast.Identifier{
											BaseNode: ast.BaseNode{
												Errors: nil,
												Loc: &ast.SourceLocation{
													End: ast.Position{
														Column: 65,
														Line:   54,
													},
													File:   "histogramQuantile_test.flux",
													Source: "loadMem",
													Start: ast.Position{
														Column: 58,
														Line:   54,
													},
												},
											},
											Name: "loadMem",
										},
									},
								},
							}, &ast.Property{
								BaseNode: ast.BaseNode{
									Errors: nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 104,
											Line:   54,
										},
										File:   "histogramQuantile_test.flux",
										Source: "fn: t_histogramQuantile",
										Start: ast.Position{
											Column: 81,
											Line:   54,
										},
									},
								},
								Key: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 83,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "fn",
											Start: ast.Position{
												Column: 81,
												Line:   54,
											},
										},
									},
									Name: "fn",
								},
								Value: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Errors: nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 104,
												Line:   54,
											},
											File:   "histogramQuantile_test.flux",
											Source: "t_histogramQuantile",
											Start: ast.Position{
												Column: 85,
												Line:   54,
											},
										},
									},
									Name: "t_histogramQuantile",
								},
							}},
							With: nil,
						},
					},
					Params: []*ast.Property{},
				},
			},
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 106,
						Line:   54,
					},
					File:   "histogramQuantile_test.flux",
					Source: "test _histogramQuantile = () => \n({input: testing.loadStorage(csv: inData), want: testing.loadMem(csv: outData), fn: t_histogramQuantile})",
					Start: ast.Position{
						Column: 1,
						Line:   53,
					},
				},
			},
		}},
		Imports: []*ast.ImportDeclaration{&ast.ImportDeclaration{
			As: nil,
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 33,
						Line:   3,
					},
					File:   "histogramQuantile_test.flux",
					Source: "import \"experimental/prometheus\"",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 33,
							Line:   3,
						},
						File:   "histogramQuantile_test.flux",
						Source: "\"experimental/prometheus\"",
						Start: ast.Position{
							Column: 8,
							Line:   3,
						},
					},
				},
				Value: "experimental/prometheus",
			},
		}, &ast.ImportDeclaration{
			As: nil,
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   4,
					},
					File:   "histogramQuantile_test.flux",
					Source: "import \"testing\"",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			Path: &ast.StringLiteral{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 17,
							Line:   4,
						},
						File:   "histogramQuantile_test.flux",
						Source: "\"testing\"",
						Start: ast.Position{
							Column: 8,
							Line:   4,
						},
					},
				},
				Value: "testing",
			},
		}},
		Metadata: "parser-type=rust",
		Name:     "histogramQuantile_test.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 24,
						Line:   1,
					},
					File:   "histogramQuantile_test.flux",
					Source: "package prometheus_test",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 24,
							Line:   1,
						},
						File:   "histogramQuantile_test.flux",
						Source: "prometheus_test",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "prometheus_test",
			},
		},
	}},
	Package: "prometheus_test",
	Path:    "testing/prometheus",
}}
