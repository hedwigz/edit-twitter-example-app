# Tools

### GraphQL Merge Tool:

1. Install deps: `npm i`
2. Compile TS: `npm run build`
3. Execute: `./bin/gql-merge` or `npx ts-node index.ts`

Flags:

```
Usage: gql-merge [options]

Options:
  -i, --input-dir <path>    input dir where you have your gql schema files (default: "../../atlantis/graphql/schema/")
  -o, --output-path <path>  output where to write the merged schema (default: "./schema.graphql")
  -h, --help                output usage information
```

[To Debug Typescript @ Jetbrains IDEs](https://www.jetbrains.com/help/idea/running-and-debugging-typescript.html)