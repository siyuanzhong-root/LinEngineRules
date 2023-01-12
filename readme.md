AST是抽象语法树（Abstract Syntax Tree）的简称，AST以树状形式表现编程语言的语法结构，树上每个节点都表示源代码中的一种结构。之所以说语法是“抽象”的，是因为这里的语法并不会表示出真实语法中出现的每个细节。

`ast`包声明了用于表示Go包的语法树的类型

## ast.Node

Golang会将所有可以识别的`token`抽象成为`ast.Node`，通过接口方式组织在一起。

AST节点实现了`ast.Node`接口，返回AST中的一个位置。

```
type Node interface {
    Pos() token.Pos // 开始位置
    End() token.Pos // 结束位置
}

```

整个AST语法树由不同的节点组成，节点可分为三种，均会实现`ast.Node`接口。

- `ast.Expr` 代表表达式和类型的节点

```
type Expr interface {
    Node
    exprNode()
}

```

- `ast.Stmt` 代表表节点

```
type Stmt interface {
    Node
    stmtNode()
}

```

- `ast.Decl` 代表声明节点

```
type Decl interface {
    Node
    declNode()
}

```

![img](http://upload-images.jianshu.io/upload_images/4933701-624604d10767c93b.png?imageMogr2/auto-orient/strip|imageView2/2/w/1056/format/webp)

ast.Node

## 示例

![img](http://upload-images.jianshu.io/upload_images/4933701-425a3123d8c09065.png?imageMogr2/auto-orient/strip|imageView2/2/w/670/format/webp)

使用ast.Node表示结构体

```
package test

import (
    "go/ast"
    "go/parser"
    "go/token"
    "testing"
)

func TestAst(t *testing.T) {
    src := `
package main
func main(){
    println("hello world")
}
    `
    //创建用于解析源文件的对象
    fileset := token.NewFileSet()
    //解析源文件，返回ast.File原始文档类型的结构体。
    astfile, err := parser.ParseFile(fileset, "", src, 0)
    if err != nil {
        panic(err)
    }
    //查看日志打印
    ast.Print(fileset, astfile)
}

```

```
0  *ast.File {
1  .  Package: 2:1
2  .  Name: *ast.Ident {
3  .  .  NamePos: 2:9
4  .  .  Name: "main"
5  .  }
6  .  Decls: []ast.Decl (len = 1) {
7  .  .  0: *ast.FuncDecl {
8  .  .  .  Name: *ast.Ident {
9  .  .  .  .  NamePos: 3:6
10  .  .  .  .  Name: "main"
11  .  .  .  .  Obj: *ast.Object {
12  .  .  .  .  .  Kind: func
13  .  .  .  .  .  Name: "main"
14  .  .  .  .  .  Decl: *(obj @ 7)
15  .  .  .  .  }
16  .  .  .  }
17  .  .  .  Type: *ast.FuncType {
18  .  .  .  .  Func: 3:1
19  .  .  .  .  Params: *ast.FieldList {
20  .  .  .  .  .  Opening: 3:10
21  .  .  .  .  .  Closing: 3:11
22  .  .  .  .  }
23  .  .  .  }
24  .  .  .  Body: *ast.BlockStmt {
25  .  .  .  .  Lbrace: 3:12
26  .  .  .  .  List: []ast.Stmt (len = 1) {
27  .  .  .  .  .  0: *ast.ExprStmt {
28  .  .  .  .  .  .  X: *ast.CallExpr {
29  .  .  .  .  .  .  .  Fun: *ast.Ident {
30  .  .  .  .  .  .  .  .  NamePos: 4:2
31  .  .  .  .  .  .  .  .  Name: "println"
32  .  .  .  .  .  .  .  }
33  .  .  .  .  .  .  .  Lparen: 4:9
34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
37  .  .  .  .  .  .  .  .  .  Kind: STRING
38  .  .  .  .  .  .  .  .  .  Value: "\"hello world\""
39  .  .  .  .  .  .  .  .  }
40  .  .  .  .  .  .  .  }
41  .  .  .  .  .  .  .  Ellipsis: -
42  .  .  .  .  .  .  .  Rparen: 4:23
43  .  .  .  .  .  .  }
44  .  .  .  .  .  }
45  .  .  .  .  }
46  .  .  .  .  Rbrace: 5:1
47  .  .  .  }
48  .  .  }
49  .  }
50  .  Scope: *ast.Scope {
51  .  .  Objects: map[string]*ast.Object (len = 1) {
52  .  .  .  "main": *(obj @ 11)
53  .  .  }
54  .  }
55  .  Unresolved: []*ast.Ident (len = 1) {
56  .  .  0: *(obj @ 29)
57  .  }
58  }

```

分析方式

按照深度优先的顺序遍历AST节点，通过递归调用`ast.Inspect()`方法来逐一打印每个节点。
如果直接打印AST则会看到一些无法被人类阅读的东西。为了防止这种情况的发生，使用`ast.Print()`来实现对AST的人工读取。

```
func TestAst(t *testing.T) {
    abspath, err := filepath.Abs("./demo.go")
    if err != nil {
        panic(abspath)
    }
    //创建用于解析源文件的对象
    fset := token.NewFileSet()
    //解析源文件，返回ast.File原始文档类型的结构体。
    f, err := parser.ParseFile(fset, abspath, nil, parser.AllErrors)
    if err != nil {
        panic(err)
    }
    //递归调用逐一打印节点
    ast.Inspect(f, func(n ast.Node) bool {
        ast.Print(fset, n)
        return true
    })
}

```

## ast.File

- `ast.File`是所有AST节点的根，仅实现`ast.Node`接口。

```
type File struct {
    Doc        *CommentGroup   // associated documentation; or nil
    Package    token.Pos       // position of "package" keyword
    Name       *Ident          // package name
    Decls      []Decl          // top-level declarations; or nil
    Scope      *Scope          // package scope (this file only)
    Imports    []*ImportSpec   // imports in this file
    Unresolved []*Ident        // unresolved identifiers in this file
    Comments   []*CommentGroup // list of all comments in the source file
}

```

- `ast.File`具有引用包名、导入声明、函数声明子节点。

![img](http://upload-images.jianshu.io/upload_images/4933701-8a2c96bb3ef5dc86.png?imageMogr2/auto-orient/strip|imageView2/2/w/705/format/webp)

ast.File

| 子节点          | 包名                 | 描述   |
| ------------ | ------------------ | ---- |
| ast.Indent   | Package Name       | 包名   |
| ast.GenDecl  | Import Declaration | 导入声明 |
| ast.FuncDecl | Func Declaration   | 函数声明 |

## ast.Indent

- 一个包名可以使用AST节点类型`*ast.Indent`来表示
- `ast.Indent`实现了`ast.Expr`接口
- 所有标识符都由`ast.Indent`结构来表示，主要包含了包的名称和在文件集中的源位置。

```
*ast.Ident {
 NamePos: dummy.go:1:9
 Name: "hello"
}

```

例如：

```
2  .  Name: *ast.Ident {
3  .  .  NamePos: 2:9
4  .  .  Name: "main"
5  .  }

```

Golang具有一个`scope`的概念，即源文件的`scope`，其中标识符表示指定的常量、类型、变量、函数、标签、包。

```
8  .  .  .  Name: *ast.Ident {
9  .  .  .  .  NamePos: 3:6
10  .  .  .  .  Name: "main"
11  .  .  .  .  Obj: *ast.Object {
12  .  .  .  .  .  Kind: func
13  .  .  .  .  .  Name: "main"
14  .  .  .  .  .  Decl: *(obj @ 7)
15  .  .  .  .  }
16  .  .  .  }

```

## ast.GenDecl

- `ast.GenDecl`表示除函数以外的所有导入声明，即`import`、`const`、`var`、`type`。
- `Tok`代表一个词性标识，用于指明声明的内容。

```
*ast.GenDecl {
 TokPos: dummy.go:3:1
 Tok: import
 Lparen: -
 Specs: []ast.Spec (len = 1) {
    0: *ast.ImportSpec {/* Omission */}
  }
  Rparen: -
}

```

## ast.ImportSpec

- 一个`ast.ImportSpec`节点对应一个导入声明。
- `ast.ImportSpec`实现了`ast.Spec`接口，访问路径可以让导入路径更有意义。

```
*ast.ImportSpec {
  Path: *ast.BasicLit {/* Omission */}
  EndPos: -
}

```

## ast.BasicLit

- 一个`ast.BasicLit`节点表示一个基本类型的文字，实现了`ast.Expr`接口。
- `ast.BasicLit`包含一个`token`类型，可使用`token.INIT`、`token.FLOAT`、`token.IMAG`、`token.CHAR`、`token.STRING`。

```
*ast.BasicLit {
  ValuePos: dummy.go:3:8
  Kind: STRING
  Value: "\"fmt\""
}

```

## ast.FuncDecl

- 一个`ast.FuncDecl`节点代表一个函数声明，仅实现了`ast.Node`接口。

## ast.FuncType

- 一个`ast.FuncType`包含一个函数签名，包括参数、结果和`func`关键字的位置。

## ast.FieldList

- `ast.FieldList`节点表示一个`Field`的咧白哦，使用括号或大括号起来。
- 列表字段是`*ast.Field`的一个切片，包含一对标识符和类型。

## ast.BlockStmt

- 一个`ast.BlockStmt`节点表示一个括号内的语句列表，实现了`ast.Stmt`接口。

## ast.ExprStmt

- `ast.ExprStmt`在语句列表中表示一个表达式，实现了`ast.Stmt`接口，并包含一个`ast.Expr`。

## ast.CallExpr

- `ast.CallExpr`表示一个调用函数的表达式，要查看的字段是：Fun、要调用的函数和Args、要传递给它的参数列表。

## ast.SelectorExpr

- `ast.SelectorExpr`表示一个带有选择器的表达式
