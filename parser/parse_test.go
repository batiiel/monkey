package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {

		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d")
		}
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		// val := stmt.(*ast.LetStatement).Value
		// if !testLetExpression(t, val, tt.expectedValue) {
		// 	return
		// }
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

// func TestReturnStatements(t *testing.T) {
// 	input := `
// 	return 5;
// 	return 10;
// 	return 83383;
// 		`
// 	l := lexer.New(input)
// 	p := New(l)

// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)
// 	if program == nil {
// 		t.Fatalf("ParseProgram() returned nil")
// 	}
// 	if len(program.Statements) != 3 {
// 		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
// 	}
// 	for _, stmt := range program.Statements {
// 		returnStms, ok := stmt.(*ast.ReturnStatement)
// 		if !ok {
// 			t.Errorf("s not *ast.ReturnStatement. got=%T", stmt)
// 			continue
// 		}
// 		if returnStms.TokenLiteral() != "return" {
// 			t.Errorf("returnStmt.TokenLiteral() not 'return', got %q", returnStms.TokenLiteral())
// 		}
// 	}
// }

// func TestIdedtifierExpression(t *testing.T) {
// 	input := "foobar;"

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	if len(program.Statements) != 1 {
// 		t.Errorf("progarm has not enough statements. got=%d", len(program.Statements))
// 	}

// 	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 	if !ok {
// 		t.Errorf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
// 	}

// 	ident, ok := stmt.Expression.(*ast.Identifier)
// 	if !ok {
// 		t.Errorf("exp not *ast.Identifier. got=%T", stmt.Expression)
// 	}

// 	if ident.Value != "foobar" {
// 		t.Errorf("ident.Value nos %s. got=%s", "foobar", ident.Value)
// 	}

// 	if ident.TokenLiteral() != "foobar" {
// 		t.Errorf("ident.TokenLiteral() not %s. got=%s", "foobar", ident.TokenLiteral())
// 	}
// }

// func TestIntegerLiteralExpression(t *testing.T) {
// 	input := "5;"

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	if len(program.Statements) != 1 {
// 		t.Fatalf("program has not enough statements. got=%d", program.Statements)
// 		return
// 	}
// 	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 	if !ok {
// 		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
// 	}

// 	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
// 	if !ok {
// 		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
// 	}

// 	if literal.Value != 5 {
// 		t.Fatalf("literal.Value not %d, got=%d", 5, literal.Value)
// 	}
// 	if literal.TokenLiteral() != "5" {
// 		t.Errorf("literal.TokenLiteral() nos %s, got=%s", "5", literal.TokenLiteral())
// 	}
// }

// func TestParsingPrefixExpression(t *testing.T) {
// 	prefixTests := []struct {
// 		input        string
// 		operator     string
// 		integerValue int64
// 	}{
// 		{"!5;", "!", 5},
// 		{"-15;", "-", 15},
// 	}

// 	for _, tt := range prefixTests {
// 		l := lexer.New(tt.input)
// 		p := New(l)
// 		program := p.ParseProgram()
// 		checkParserErrors(t, p)

// 		if len(program.Statements) != 1 {
// 			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
// 		}
// 		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 		if !ok {
// 			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
// 		}

// 		exp, ok := stmt.Expression.(*ast.PrefixExpression)
// 		if !ok {
// 			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Expression)
// 		}

// 		if exp.Operator != tt.operator {
// 			t.Fatalf("exp.Operator not '%s', got=%s", tt.operator, exp.Operator)
// 		}
// 		if !testIntegerLiteral(t, exp.Rigth, tt.integerValue) {
// 			return
// 		}
// 	}
// }

// func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
// 	integ, ok := il.(*ast.IntegerLiteral)
// 	if !ok {
// 		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
// 		return false
// 	}
// 	if integ.Value != value {
// 		t.Errorf("integ.Valu not %d. got=%d", value, integ.Value)
// 		return false
// 	}
// 	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
// 		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
// 		return false
// 	}
// 	return true
// }

// func TestParsingInfixExpressions(t *testing.T) {
// 	infixTests := []struct {
// 		input      string
// 		leftValue  int64
// 		operator   string
// 		rightValue int64
// 	}{
// 		{"5 + 5;", 5, "+", 5},
// 		{"5 - 5;", 5, "-", 5},
// 		{"5 * 5;", 5, "*", 5},
// 		{"5 / 5;", 5, "/", 5},
// 		{"5 > 5;", 5, ">", 5},
// 		{"5 < 5;", 5, "<", 5},
// 		{"5 == 5;", 5, "==", 5},
// 		{"5 != 5;", 5, "!=", 5},
// 	}

// 	for _, tt := range infixTests {
// 		l := lexer.New(tt.input)
// 		p := New(l)
// 		program := p.ParseProgram()
// 		checkParserErrors(t, p)

// 		if len(program.Statements) != 1 {
// 			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
// 				1, len(program.Statements))
// 		}

// 		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 		if !ok {
// 			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
// 				program.Statements[0])
// 		}

// 		exp, ok := stmt.Expression.(*ast.InfixExpression)
// 		if !ok {
// 			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
// 		}

// 		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
// 		}

// 		if exp.Operator != tt.operator {
// 			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
// 		}

// 		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
// 		}
// 	}
// }

// func TestOperatorPrecedenceParsing(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{
// 			"-a * b",
// 			"((-a) * b)",
// 		},
// 		{
// 			"!-a",
// 			"(!(-a))",
// 		},
// 		{
// 			"a + b + c",
// 			"((a + b) + c)",
// 		},
// 		{
// 			"a + b - c",
// 			"((a + b) - c)",
// 		},
// 		{
// 			"a * b * c",
// 			"((a * b) * c)",
// 		},
// 		{
// 			"a * b / c",
// 			"((a * b) / c)",
// 		},
// 		{
// 			"a + b / c",
// 			"(a + (b / c))",
// 		},
// 		{
// 			"a + b * c + d / e - f",
// 			"(((a + (b * c)) + (d / e)) - f)",
// 		},
// 		{
// 			"3 + 4; -5 * 5",
// 			"(3 + 4)((-5) * 5)",
// 		},
// 		{
// 			"5 > 4 == 3 < 4",
// 			"((5 > 4) == (3 < 4))",
// 		},
// 		{
// 			"5 < 4 != 3 > 4",
// 			"((5 < 4) != (3 > 4))",
// 		},
// 		{
// 			"3 + 4 * 5 == 3 * 1 + 4 * 5",
// 			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
// 		},
// 		{
// 			"true",
// 			"true",
// 		},
// 		{
// 			"false",
// 			"false",
// 		},
// 		{
// 			"3 > 5 == false",
// 			"((3 > 5) == false)",
// 		},
// 		{
// 			"3 < 5 == true",
// 			"((3 < 5) == true)",
// 		},
// 		{
// 			"1 + (2 + 3) + 4",
// 			"((1 + (2 + 3)) + 4)",
// 		},
// 		{
// 			"(5 + 5) * 2",
// 			"((5 + 5) * 2)",
// 		},
// 		{
// 			"2 / (5 + 5)",
// 			"(2 / (5 + 5))",
// 		},
// 		{
// 			"(5 + 5) * 2 * (5 + 5)",
// 			"(((5 + 5) * 2) * (5 + 5))",
// 		},
// 		{
// 			"-(5 + 5)",
// 			"(-(5 + 5))",
// 		},
// 		{
// 			"!(true == true)",
// 			"(!(true == true))",
// 		},
// 		{
// 			"a + add(b * c) + d",
// 			"((a + add((b * c))) + d)",
// 		},
// 		{
// 			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
// 			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
// 		},
// 		{
// 			"add(a + b + c * d / f + g)",
// 			"add((((a + b) + ((c * d) / f)) + g))",
// 		},
// 	}

// 	for _, tt := range tests {
// 		l := lexer.New(tt.input)
// 		p := New(l)
// 		program := p.ParseProgram()
// 		checkParserErrors(t, p)

// 		actual := program.String()
// 		if actual != tt.expected {
// 			t.Errorf("expected=%q, got=%q", tt.expected, actual)
// 		}
// 	}
// }

// func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
// 	ident, ok := exp.(*ast.Identifier)
// 	if !ok {
// 		t.Errorf("exp not *ast.Identifier. got=%T", exp)
// 		return false
// 	}

// 	if ident.Value != value {
// 		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
// 		return false
// 	}

// 	if ident.TokenLiteral() != value {
// 		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
// 		return false
// 	}

// 	return true

// }

// func testLiteralExpression(
// 	t *testing.T,
// 	exp ast.Expression,
// 	expected interface{},
// ) bool {
// 	switch v := expected.(type) {
// 	case int:
// 		return testIntegerLiteral(t, exp, int64(v))
// 	case int64:
// 		return testIntegerLiteral(t, exp, v)
// 	case string:
// 		return testIdentifier(t, exp, v)
// 	case bool:
// 		return testBooleanLiteral(t, exp, v)
// 	}
// 	t.Errorf("type of exp not handled. got=%T", exp)
// 	return false
// }

// func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
// 	operator string, right interface{}) bool {

// 	opExp, ok := exp.(*ast.InfixExpression)
// 	if !ok {
// 		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
// 		return false
// 	}

// 	if !testLiteralExpression(t, opExp.Left, left) {
// 		return false
// 	}

// 	if opExp.Operator != operator {
// 		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
// 		return false
// 	}

// 	if !testLiteralExpression(t, opExp.Right, right) {
// 		return false
// 	}

// 	return true
// }

// func testBooleanLiteral(
// 	t *testing.T,
// 	exp ast.Expression,
// 	value bool,
// ) bool {
// 	bo, ok := exp.(*ast.Boolean)
// 	if !ok {
// 		t.Errorf("exp not *ast.Boolean. got=%T", exp)
// 		return false
// 	}

// 	if bo.Value != value {
// 		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
// 		return false
// 	}

// 	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
// 		t.Errorf("bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
// 		return false
// 	}

// 	return true
// }
