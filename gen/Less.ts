
import * as ts from "typescript";
import * as fs from "fs";
import * as path from "path";


export enum FieldType {
    INT32,
    INT64,
    FLOAT32,
    FLOAT64,
    STRING,
    BOOLEAN,
    OBJECT,
    FILE,
    ENUM
}

export interface LessTag {
    name: string
    text?: string
}


export interface LessEnumItem {
    name: string
    value: number | string
    title: string
}

export enum LessEnumType {
    INT32,
    STRING,
}
export interface LessEnum {
    title: string
    name: string
    path: string
    items: LessEnumItem[]
    tags: LessTag[]
    type: LessEnumType
}

export interface LessObject {
    title: string
    name: string
    path: string
    fields: LessField[]
    tags: LessTag[]
}

export interface LessField {
    type: FieldType
    name: string
    title: string
    required: boolean
    typeSymbol?: string
    isArray: boolean
    pattern?: string
    output?: boolean
    length?: number
    index?: string
    unique?: string
    tags: LessTag[]
}

export interface LessRequest {
    title: string
    method: string
    fields: LessField[]
    tags: LessTag[]
}

export interface LessResponse {
    title: string
    fields: LessField[]
    tags: LessTag[]
}

export interface Less {
    name: string
    request: LessRequest
    response: LessResponse
    enums: LessEnum[]
    objects: LessObject[]
}

function walkFile(p: string, files: string[]): void {
    if (p.endsWith(".ts")) {
        files.push(p);
    } else {
        let ts = fs.statSync(p);
        if (ts && ts.isDirectory()) {
            let items = fs.readdirSync(p);
            for (let item of items) {
                if (item.startsWith(".")) {
                    continue;
                }
                walkFile(path.join(p, item), files);
            }
        }
    }
}

interface KeySet {
    [key: string]: boolean
}

function lessFieldType(node: ts.TypeNode | undefined, program: ts.Program, field: LessField, createSymbol: (symbol: ts.Symbol) => void): void {

    if (node === undefined) {
        return;
    }

    let checker = program.getTypeChecker();

    let type = checker.getTypeAtLocation(node);

    if (type.aliasSymbol !== undefined) {
        switch (type.aliasSymbol.name) {
            case "int32":
                field.type = FieldType.INT32
                return;
            case "int64":
                field.type = FieldType.INT64
                return;
            case "float32":
                field.type = FieldType.FLOAT32
                return;
            case "float64":
                field.type = FieldType.FLOAT64
                return;
        }
    }

    if (type.flags & (ts.TypeFlags.Enum | ts.TypeFlags.EnumLike | ts.TypeFlags.EnumLiteral)) {
        if (type.symbol !== undefined) {
            field.type = FieldType.ENUM
            field.typeSymbol = type.symbol.name;
            createSymbol(type.symbol);
        }
    } else if (type.flags & (ts.TypeFlags.Number | ts.TypeFlags.NumberLike | ts.TypeFlags.NumberLiteral)) {
        field.type = FieldType.FLOAT64
        switch (node.getText()) {
            case "int32":
                field.type = FieldType.INT32
                return;
            case "int64":
                field.type = FieldType.INT64
                return;
            case "float32":
                field.type = FieldType.FLOAT32
                return;
            case "float64":
                field.type = FieldType.FLOAT64
                return;
        }
        return;
    } else if (type.flags & (ts.TypeFlags.Boolean | ts.TypeFlags.BooleanLike | ts.TypeFlags.BooleanLiteral)) {
        field.type = FieldType.BOOLEAN
        return;
    } else if (type.flags & (ts.TypeFlags.String | ts.TypeFlags.StringLike | ts.TypeFlags.StringLiteral)) {
        field.type = FieldType.STRING
        return;
    } else if (type.flags & ts.TypeFlags.Object) {
        if (type.symbol !== undefined) {
            if (type.symbol.name == "Array") {
                ts.forEachChild(node, (s: ts.Node): void => {
                    if (ts.isTypeNode(s)) {
                        lessFieldType(s, program, field, createSymbol);
                    }
                });
                field.isArray = true;
            } else {
                field.type = FieldType.OBJECT
                field.typeSymbol = type.symbol.name
                createSymbol(type.symbol);
            }
        } else {
            field.type = FieldType.OBJECT
        }
        return;

    } else if (type.isUnion()) {

        let vs: ts.TypeNode[] = [];

        ts.forEachChild(node, (s: ts.Node): void => {
            if (ts.isTypeNode(s)) {
                vs.push(s);
            }
        });

        vs = vs.reverse();

        for (let s of vs) {

            lessFieldType(s, program, field, createSymbol);
        }
    } else if (type.flags & ts.TypeFlags.Any) {
        field.type = FieldType.OBJECT
        field.typeSymbol = undefined;
    }

    return;
}

function lessField(node: ts.InterfaceDeclaration | ts.ClassDeclaration, program: ts.Program, fields: LessField[], keySet: KeySet, createSymbol: (symbol: ts.Symbol) => void): void {

    let checker = program.getTypeChecker();

    if (node.members !== undefined) {

        for (let member of node.members) {

            if (ts.isPropertyDeclaration(member) || ts.isPropertySignature(member)) {

                let symbol = checker.getSymbolAtLocation(member.name!)!;

                if (keySet[symbol.name]) {
                    continue;
                }

                keySet[symbol.name] = true;

                let title = ts.displayPartsToString(symbol.getDocumentationComment(checker)) || '';
                let field: LessField = {
                    title: title,
                    name: symbol.name,
                    type: FieldType.OBJECT,
                    required: member.questionToken === undefined,
                    isArray: false,
                    tags: []
                };

                for (let tag of symbol.getJsDocTags()) {
                    if (tag.name == "pattern" && tag.text !== undefined) {
                        field.pattern = tag.text
                    }
                    if (tag.name == "output" && tag.text !== undefined) {
                        field.output = tag.text != 'false'
                    }
                    if (tag.name == "length" && tag.text !== undefined) {
                        field.length = parseInt(tag.text);
                    }
                    if (tag.name == "index" && tag.text !== undefined) {
                        field.index = tag.text;
                    }
                    if (tag.name == "unique" && tag.text !== undefined) {
                        field.unique = tag.text;
                    }
                    field.tags.push({ name: tag.name, text: tag.text });
                }

                lessFieldType(member.type, program, field, createSymbol);

                fields.push(field);
            }

        }
    }
}

function lessRequest(node: ts.InterfaceDeclaration | ts.ClassDeclaration, program: ts.Program, createSymbol: (symbol: ts.Symbol) => void): LessRequest | undefined {

    let checker = program.getTypeChecker();
    let symbol = checker.getSymbolAtLocation(node.name!)!;

    let title = ts.displayPartsToString(symbol.getDocumentationComment(checker)) || '';
    let method = "GET";
    let tags: LessTag[] = [];

    for (let tag of symbol.getJsDocTags()) {
        if (tag.name == "method" && tag.text) {
            method = tag.text;
        }
        tags.push({ name: tag.name, text: tag.text })
    }

    let keySet: KeySet = {};
    let fields: LessField[] = [];


    if (node.heritageClauses !== undefined) {
        for (let clause of node.heritageClauses) {
            for (let node of clause.types) {
                let type = checker.getTypeAtLocation(node);
                for (let declaration of type.symbol.declarations) {
                    if (ts.isInterfaceDeclaration(declaration) || ts.isClassDeclaration(declaration)) {
                        lessField(declaration, program, fields, keySet, createSymbol);
                    }
                }
            }
        }
    }

    lessField(node, program, fields, keySet, createSymbol);

    return {
        title: title,
        method: method,
        fields: fields,
        tags: tags
    };
}

function lessResponse(node: ts.InterfaceDeclaration | ts.ClassDeclaration, program: ts.Program, createSymbol: (symbol: ts.Symbol) => void): LessResponse | undefined {

    let checker = program.getTypeChecker();
    let symbol = checker.getSymbolAtLocation(node.name!)!;
    let tags: LessTag[] = [];

    let title = ts.displayPartsToString(symbol.getDocumentationComment(checker)) || '';

    for (let tag of symbol.getJsDocTags()) {
        tags.push({ name: tag.name, text: tag.text })
    }

    let keySet: KeySet = {};
    let fields: LessField[] = [];

    if (node.heritageClauses !== undefined) {
        for (let clause of node.heritageClauses) {
            for (let node of clause.types) {
                let type = checker.getTypeAtLocation(node);
                for (let declaration of type.symbol.declarations) {
                    if (ts.isInterfaceDeclaration(declaration) || ts.isClassDeclaration(declaration)) {
                        lessField(declaration, program, fields, keySet, createSymbol);
                    }
                }
            }
        }
    }

    lessField(node, program, fields, keySet, createSymbol);

    return {
        title: title,
        fields: fields,
        tags: tags
    };
}

function lessObject(symbol: ts.Symbol, node: ts.ClassDeclaration | ts.InterfaceDeclaration, program: ts.Program, createSymbol: (symbol: ts.Symbol) => void): LessObject | undefined {

    let checker = program.getTypeChecker();

    let title = ts.displayPartsToString(symbol.getDocumentationComment(checker)) || '';
    let tags: LessTag[] = [];

    for (let tag of symbol.getJsDocTags()) {
        tags.push({ name: tag.name, text: tag.text })
    }

    let keySet: KeySet = {};
    let fields: LessField[] = [];

    if (node.heritageClauses !== undefined) {
        for (let clause of node.heritageClauses) {
            for (let node of clause.types) {
                let type = checker.getTypeAtLocation(node);
                for (let declaration of type.symbol.declarations) {
                    if (ts.isInterfaceDeclaration(declaration) || ts.isClassDeclaration(declaration)) {
                        lessField(declaration, program, fields, keySet, createSymbol);
                    }
                }
            }
        }
    }

    lessField(node, program, fields, keySet, createSymbol);

    return {
        title: title,
        name: symbol.name,
        path: node.getSourceFile().fileName,
        fields: fields,
        tags: tags
    };
}

function lessEnum(symbol: ts.Symbol, node: ts.EnumDeclaration, program: ts.Program, createSymbol: (symbol: ts.Symbol) => void): LessEnum | undefined {

    let checker = program.getTypeChecker();

    let items: LessEnumItem[] = [];
    let title = ts.displayPartsToString(symbol.getDocumentationComment(checker)) || '';

    let tags: LessTag[] = [];

    let type = LessEnumType.INT32;

    for (let tag of symbol.getJsDocTags()) {
        tags.push({ name: tag.name, text: tag.text })
    }

    for (let member of node.members) {

        let s = checker.getSymbolAtLocation(member.name);

        let item = {
            name: member.name.getText(),
            value: member.initializer === undefined ? items.length : eval(member.initializer.getText()),
            title: s === undefined ? '' : ts.displayPartsToString(s.getDocumentationComment(checker))
        };

        items.push(item)

        if (typeof item.value == 'string') {
            type = LessEnumType.STRING
        }
    }

    return {
        title: title,
        name: symbol.name,
        path: node.getSourceFile().fileName,
        items: items,
        tags: tags,
        type: type
    };
}


function lessFile(basePath: string, file: ts.SourceFile, program: ts.Program): Less | undefined {

    let name = path.relative(basePath, file.fileName);

    if (name.endsWith(".less.ts")) {
        name = name.substr(0, name.length - 8);
    }

    let request: LessRequest | undefined
    let response: LessResponse | undefined;
    let symbolKey: KeySet = {};
    let enums: LessEnum[] = [];
    let objects: LessObject[] = [];

    function createSymbol(symbol: ts.Symbol): void {

        if (symbolKey[symbol.name]) {
            return;
        }

        symbolKey[symbol.name] = true;

        if (symbol.declarations !== undefined) {

            for (let decl of symbol.declarations) {

                if (ts.isEnumDeclaration(decl)) {
                    let v = lessEnum(symbol, decl, program, createSymbol);
                    if (v !== undefined) {
                        enums.push(v);
                    }
                } else if (ts.isClassDeclaration(decl) || ts.isInterfaceDeclaration(decl)) {
                    let v = lessObject(symbol, decl, program, createSymbol);
                    if (v !== undefined) {
                        objects.push(v);
                    }
                }

                break;
            }
        }

    }

    let checker = program.getTypeChecker();

    function each(node: ts.Node): void {
        if (ts.isInterfaceDeclaration(node) || ts.isClassDeclaration(node)) {
            let name = checker.getSymbolAtLocation(node.name!)!.name;
            if (name == "Request") {
                request = lessRequest(node, program, createSymbol);
                return;
            } else if (name == "Response") {
                response = lessResponse(node, program, createSymbol);
                return;
            }
        } if (ts.isClassDeclaration(node)) {
            lessObject(checker.getSymbolAtLocation(node.name!)!, node, program, createSymbol)
        } if (ts.isEnumDeclaration(node)) {
            lessEnum(checker.getSymbolAtLocation(node.name!)!, node, program, createSymbol)
        }
    }

    ts.forEachChild(file, each);


    if (request !== undefined && response !== undefined) {
        return {
            name: name,
            request: request,
            response: response,
            enums: enums.reverse(),
            objects: objects.reverse()
        }
    }

    return undefined;
}

export function walk(basePath: string, cb: (less: Less) => void): void {

    let files: string[] = [];

    walkFile(basePath, files);

    let program = ts.createProgram({
        rootNames: files,
        options: {
            target: ts.ScriptTarget.ES5,
            module: ts.ModuleKind.CommonJS,
            removeComments: false
        }
    })

    for (let sourceFile of program.getSourceFiles()) {

        if (sourceFile.fileName.endsWith(".less.ts")) {
            let less = lessFile(basePath, sourceFile, program);
            if (less !== undefined) {
                cb(less);
            }
        }
    }

}