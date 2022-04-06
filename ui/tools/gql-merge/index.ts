import { promises as fsPromises } from "fs";
import * as path from "path";
import { Command } from "commander";
import { mergeTypeDefs } from "@graphql-tools/merge";
import { print } from "graphql";

const program = new Command();

program
  .option(
    "-i, --input-dir <path>",
    "input dir where you have your gql schema files",
    "../../../atlantis/graphql/schema/"
  )
  .option(
    "-o, --output-path <path>",
    "output where to write the merged schema",
    "./schema.graphql"
  )
  .parse(process.argv);

(async function getSchemas() {
  try {
    const options = program.opts();
    const schemaFiles = (await fsPromises.readdir(options.inputDir)).filter(d => d.includes(".graphql"));
    console.log("hedwigz files to merge:", schemaFiles);
    const schemas: string[] = [];
    for (let file of schemaFiles) {
      file = path.join(options.inputDir, file);
      schemas.push(await fsPromises.readFile(file, { encoding: "utf8" }));
    }
    const merged = mergeTypeDefs(schemas);
    await fsPromises.writeFile(options.outputPath, print(merged));
    console.log(`files merged to ${options.outputPath}`);
  } catch (e) {
    console.error(`error while merging schemas: ${e}`, e.stackTrace);
  }
})();
