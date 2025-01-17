class App{
    constructor(){
        self.app = this
        this.cache = {}
        this.cache.grid = {}
        this.cache.filter = {}
    }

    async init(){
        this.browser = await this.getBrowser()
    }

    // grid cache
    async getGridCache(id) {
        return this.cache.grid[id];
    }

    async setGridCache(id, data){
        this.cache.grid[id] = data
    }

    // filter cache
    async getFilterCache(id){
        return this.cache.filter[id]
    }

    async setFilterCache(id, data){
        this.cache.filter[id] = data
    }

    // browser
    async getBrowser(){
        return {
            userAgent: navigator.userAgent,
            platform: navigator.platform,
            isSupportWebp: (() => {
                const elem = document.createElement('canvas');
                if (!!(elem.getContext && elem.getContext('2d'))) {
                    return elem.toDataURL('image/webp').indexOf('data:image/webp') == 0;
                }
                return false;
            })(),
        }
    }
}

new App().init()