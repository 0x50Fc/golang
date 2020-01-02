"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const less = __importStar(require("./Less"));
const swagger = __importStar(require("./Swagger"));
const aliSwagger = __importStar(require("./ali/swagger"));
const postman = __importStar(require("./Postman"));
const golangCli = __importStar(require("./golang/cli"));
const golangServer = __importStar(require("./golang/server"));
const process = __importStar(require("process"));
const sql = __importStar(require("./sql"));
const cli = __importStar(require("./cli"));
const srv = __importStar(require("./srv"));
const srvlib = __importStar(require("./srv-lib"));
const mindmap = __importStar(require("./mindmap/less"));
const fs = __importStar(require("fs"));
const phpCli = __importStar(require("./php/cli"));
let commander = require('commander');
commander
    .command('less <inDir>')
    .description('quick generate less ')
    .usage("less <inDir>")
    .action(function (inDir) {
    less.walk(inDir, (v) => {
        process.stdout.write(JSON.stringify(v, undefined, 4));
        process.stdout.write("\n");
    });
});
commander
    .command('swagger <inDir> <baseURL> [title] [version]')
    .description('quick generate swagger ')
    .usage("swagger <inDir> <baseURL> [title] [version]")
    .action(function (inDir, baseURL, title, version) {
    let v = swagger.walk(inDir, baseURL, title, version);
    process.stdout.write(JSON.stringify(v, undefined, 4));
    process.stdout.write("\n");
});
commander
    .command('ali-swagger <inDir> <outFile> <vpc> [limit]')
    .description('quick generate ali-swagger ')
    .usage("ali-swagger <inDir> <outFile> <vpc> [limit]")
    .action(function (inDir, outFile, vpc, limit) {
    let v = aliSwagger.walk(inDir, vpc);
    if (limit === undefined) {
        fs.writeFileSync(outFile + ".swagger", JSON.stringify(v, undefined, 4));
    }
    else {
        let paths = v.paths;
        let keys = [];
        for (let key in v.paths) {
            keys.push(key);
        }
        keys.sort();
        let l = parseInt(limit) || 10;
        let n = 0;
        let p = 1;
        let ps = {};
        for (let key of keys) {
            ps[key] = paths[key];
            n++;
            if (n >= l) {
                v.paths = ps;
                fs.writeFileSync(outFile + '-' + p + ".swagger", JSON.stringify(v, undefined, 4));
                ps = {};
                n = 0;
                p++;
            }
        }
        if (n > 0) {
            v.paths = ps;
            fs.writeFileSync(outFile + '-' + p + ".swagger", JSON.stringify(v, undefined, 4));
            ps = {};
            n = 0;
            p++;
        }
    }
});
commander
    .command('postman <inDir> <baseURL> [title] [version]')
    .description('quick generate postman ')
    .usage("postman <inDir> <baseURL> [title] [version]")
    .action(function (inDir, baseURL, title, version) {
    let v = postman.walk(inDir, baseURL, title, version);
    process.stdout.write(JSON.stringify(v, undefined, 4));
    process.stdout.write("\n");
});
commander
    .command('golang <inDir> <outDir>')
    .description('quick generate golang ')
    .usage("golang <inDir> <outDir>")
    .action(function (inDir, outDir) {
    golangServer.walk(inDir, outDir);
});
commander
    .command('golang-cli <inDir> <outDir>')
    .description('quick generate golang-cli ')
    .usage("golang-cli <inDir> <outDir>")
    .action(function (inDir, outDir) {
    golangCli.walk(inDir, outDir);
});
commander
    .command('sql <inDir> <dest> <configFile>')
    .description('quick generate sql ')
    .usage("sql <inDir> <dest> <configFile>")
    .action(function (inDir, dest, configFile) {
    var config = {};
    if (fs.existsSync(configFile)) {
        let v = fs.readFileSync(configFile).toString();
        try {
            config = JSON.parse(v);
        }
        catch (e) {
            console.info("[JSON]", e);
        }
    }
    let vs = [];
    let tableSet = sql.walk(inDir, config.prefix || '', (v) => {
        vs.push(v);
    }, config.tableSet);
    fs.writeFileSync(dest + ".sql", vs.join(''), { encoding: 'utf-8' });
    fs.writeFileSync(dest + '.json', JSON.stringify({
        prefix: config.prefix || '',
        autoIncrement: config.autoIncrement,
        tableSplitSet: config.tableSplitSet,
        tableSet: tableSet
    }, undefined, 4), { encoding: 'utf-8' });
});
commander
    .command('cli <inDir> <outDir>')
    .description('quick generate cli ')
    .usage("cli <inDir> <outDir>")
    .action(function (inDir, outDir) {
    cli.walk(inDir, outDir);
});
commander
    .command('srv <inDir> <outDir>')
    .description('quick generate srv ')
    .usage("srv <inDir> <outDir>")
    .action(function (inDir, outDir) {
    srv.walk(inDir, outDir);
});
commander
    .command('srv-lib <inDir> <outDir>')
    .description('quick generate srv-lib ')
    .usage("srv-lib <inDir> <outDir>")
    .action(function (inDir, outDir) {
    srvlib.walk(inDir, outDir);
});
commander
    .command('mindmap <inDir> <outFile> <name> [title]')
    .description('quick generate mindmap ')
    .usage("mindmap <inDir> <outFile> <name> [title]")
    .action(function (inDir, outFile, name, title) {
    mindmap.walk(inDir, outFile, name, title);
});
commander
    .command('php-cli <inDir> <outFile> <ns>')
    .description('quick generate php cli ')
    .usage("php-cli <inDir> <outFile> <namespace>")
    .action(function (inDir, outFile, ns) {
    phpCli.walk(inDir, outFile, ns);
});
commander.parse(process.argv);
