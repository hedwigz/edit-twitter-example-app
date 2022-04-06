"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
var fs_1 = require("fs");
var path = require("path");
var commander_1 = require("commander");
var merge_1 = require("@graphql-tools/merge");
var graphql_1 = require("graphql");
var program = new commander_1.Command();
program
    .option("-i, --input-dir <path>", "input dir where you have your gql schema files", "../../../atlantis/graphql/schema/")
    .option("-o, --output-path <path>", "output where to write the merged schema", "./schema.graphql")
    .parse(process.argv);
(function getSchemas() {
    return __awaiter(this, void 0, void 0, function () {
        var options, schemaFiles, schemas, _i, schemaFiles_1, file, _a, _b, merged, e_1;
        return __generator(this, function (_c) {
            switch (_c.label) {
                case 0:
                    _c.trys.push([0, 7, , 8]);
                    options = program.opts();
                    return [4 /*yield*/, fs_1.promises.readdir(options.inputDir)];
                case 1:
                    schemaFiles = (_c.sent()).filter(function (d) { return d.includes(".graphql"); });
                    console.log("hedwigz files to merge:", schemaFiles);
                    schemas = [];
                    _i = 0, schemaFiles_1 = schemaFiles;
                    _c.label = 2;
                case 2:
                    if (!(_i < schemaFiles_1.length)) return [3 /*break*/, 5];
                    file = schemaFiles_1[_i];
                    file = path.join(options.inputDir, file);
                    _b = (_a = schemas).push;
                    return [4 /*yield*/, fs_1.promises.readFile(file, { encoding: "utf8" })];
                case 3:
                    _b.apply(_a, [_c.sent()]);
                    _c.label = 4;
                case 4:
                    _i++;
                    return [3 /*break*/, 2];
                case 5:
                    merged = merge_1.mergeTypeDefs(schemas);
                    return [4 /*yield*/, fs_1.promises.writeFile(options.outputPath, graphql_1.print(merged))];
                case 6:
                    _c.sent();
                    console.log("files merged to " + options.outputPath);
                    return [3 /*break*/, 8];
                case 7:
                    e_1 = _c.sent();
                    console.error("error while merging schemas: " + e_1, e_1.stackTrace);
                    return [3 /*break*/, 8];
                case 8: return [2 /*return*/];
            }
        });
    });
})();
//# sourceMappingURL=index.js.map