class WebAssemblyApp {
    constructor(appFileName) {
        self.wa = this;
        this._folderPath = "/assets/wasm/";
        this._appFileName = appFileName;
        this.appFileVersion = "1.243";
        this._go = new Go();
        this._mod = null;
        this.memory = null;
        this._inst = null;
        this._importObject = {
            wasi_snapshot_preview1: {
                proc_exit: (code) => console.log(`proc_exit called with code: ${code}`),
                ...this._go.importObject.wasi_snapshot_preview1,
            },
            ...this._go.importObject,
        };

        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }
    }

    async init() {
        try {
            const wasmPath = `${this._folderPath}${this._appFileName}?v=${this.appFileVersion}`;
            const result = await WebAssembly.instantiateStreaming(fetch(wasmPath), this._importObject);
            this._mod = result.module;
            this._inst = result.instance;
            this._func = this._inst.exports;
            await this._go.run(this._inst);
            this.memory = new Uint8Array(this._func.memory.buffer);
        } catch (err) {
            console.error("Failed to initialize WebAssembly module:", err);

        }
    }

    async #send(value = null){
        const id = Math.floor(Math.random() * (100000 - 10000 + 1)) + 10000;
        // check it is unique
        if (sessionStorage.getItem(id.toString())) {
            return this.#send(value);
        }

        // save it
        sessionStorage.setItem(id.toString(), JSON.stringify(value));
        
        return id
    }

    async #get(id){
        const v =  JSON.parse(sessionStorage.getItem(id));
        sessionStorage.removeItem(id);
        return v
    }

    // test func get status
    async status() {
        const id = await this.#send();
        this._func.status(id);
        return this.#get(id);
    }

    // grid
    async grid(request) {
        const id = await this.#send(request);
        await this._func.grid(id);
        return this.#get(id);
    }
}

// Usage Example
(async () => {
    const wasmApp = new WebAssemblyApp("adminApp.wasm");
    await wasmApp.init();

    // status
    console.log((await wasmApp.status()).status);

    // grid
    console.log(await wasmApp.grid({
        id: "123",
        sku: "sku",
    }));
})();
