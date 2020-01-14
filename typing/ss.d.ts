

declare interface HeaderSet {
    [key: string]: string
}

declare interface CookieSet {
    [key: string]: string
}

declare interface QueryObjectSet {
    [key: string]: string
}

declare interface Input {
    readonly method: string
    readonly path: string
    readonly protocol: string
    readonly query: string
    readonly host: string
    readonly data: any
    readonly clientIp: string
    readonly sessionId: string
    readonly header: HeaderSet
    readonly cookie: CookieSet
    readonly trace: string
    readonly queryObject: QueryObjectSet
}

declare var input: Input

declare interface Output {
    data: any
    readonly header: HeaderSet
    readonly cookie: CookieSet
    status: number
}

declare var output: Output

declare enum HttpType {
    Urlencode = "application/x-www-form-urlencoded",
    Json = "application/json",
    Text = "text/plain",
    Xml = "text/xml",
    FormData = "multipart/form-data"
}

declare enum HttpResponseType {
    Text = "text",
    Json = "json",
    Byte = "byte",
    Response = "response"
}

declare interface HttpOption {
    readonly url: string
    readonly method: string
    readonly type?: HttpType
    readonly responseType?: HttpResponseType
    readonly data?: any
    readonly header?: HeaderSet
    readonly timeout?: number
    readonly charset?: string
}

declare interface Http {
    send(option: HttpOption): any
}

declare var http: Http

declare interface Redis {
    readonly prefix: string
    get(key: string): string
    set(key: string, value: string, expires?: number): void
    del(key: string): void
    eval(script: string, keys: string[], args: any[]): any
}

declare var redis: Redis

declare interface Assets {
    getString(path: string): string | undefined
    get(path: string): ArrayBuffer | undefined
    has(path: string): number | undefined
}

declare var assets: Assets

declare interface FS {
    putContent(path: string, content: string | ArrayBuffer | ArrayBufferView): void
    remove(path: string): void
}

declare var fs: FS

declare var app: any

declare var __dirname: string

declare function require(path: string): any

declare interface Console {
    info(...args: any[]): void
    error(...args: any[]): void
    warn(...args: any[]): void
    debug(...args: any[]): void
}

declare var console: Console

declare interface DatabaseResult {
    id: number | string
    count: number
}

declare interface DatabaseRow {
    [name: string]: string
}

declare interface Database {
    query(name: string, sql: string, ...args: string[]): DatabaseRow[] | undefined
    exec(name: string, sql: string, ...args: string[]): DatabaseResult | undefined
    table(name: string, sql: string, ...args: string[]): string[][] | undefined
    copy(dst: string, table: string, uniqueKeys: string[], name: string, sql: string, ...args: string[]): void
    getErrmsg(): string
}

declare var db: Database

declare type int64 = number | string

declare interface Int64Interface {
    add(a: int64, b: int64): int64
    sub(a: int64, b: int64): int64
    mul(a: int64, b: int64): int64
    div(a: int64, b: int64): int64
    mod(a: int64, b: int64): int64
    comp(a: int64, b: int64): number
}

declare var Int64: Int64Interface


declare type ViewId = int64
declare type ElementId = int64

declare interface View {
    open(): ViewId
    close(id: ViewId): void
    create(id: ViewId, name: string): ElementId
    del(id: ViewId, elementId: ElementId): void
    add(id: ViewId, elementId: ElementId, name: string): ElementId
    set(id: ViewId, elementId: ElementId, key: string, value?: any): void
    setUnit(id: ViewId, name: string, scale: number, base: number): void
    toPNGData(id: ViewId, elementId: ElementId, width: number, height: number): ArrayBuffer | ArrayBufferView
    toJPGData(id: ViewId, elementId: ElementId, width: number, height: number, quality?: number): ArrayBuffer | ArrayBufferView
}

declare var view: View

declare interface MQ {
    send(conn: string, name: string, data: any): void
}

declare var mq: MQ


declare interface GeoipNameSet {
    [key: string]: string
}

declare interface GeoipCity {
    geoname_id: number
    names: GeoipNameSet
}

declare interface GeoipContinent {
    code: string
    geoname_id: number
    names: GeoipNameSet
}

declare interface GeoipCountry {
    iso_code: string
    geoname_id: number
    names: GeoipNameSet
}

declare interface GeoipLocation {
    accuracy_radius: number
    latitude: number
    longitude: number
    metro_code: number
    time_zone: string
}

declare interface GeoipCityObject {
    city: GeoipCity
    continent: GeoipContinent
    country: GeoipCountry
    location: GeoipLocation
}

declare interface GeoipCountryObject {
    continent: GeoipContinent
    country: GeoipCountry
}

declare interface Geoip {
    city(ip: string): GeoipCityObject
    country(ip: string): GeoipCountryObject
}

declare var geoip: Geoip
