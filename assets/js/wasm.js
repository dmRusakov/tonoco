const folderPath = "/assets/wasm/";
const appFileName = "adminApp.wasm";
const appFileVersion = "1.007";

if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch(folderPath + appFileName + "?v=" + appFileVersion), go.importObject).then( async (result) => {
    mod = result.module;
    inst = result.instance;
    await run();
    document.getElementById("runButton").disabled = false;
}).catch((err) => {
    console.error(err);
});

async function run() {
    // console.clear();
    await go.run(inst);
    inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}