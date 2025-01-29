class User {
    constructor() {
        self.user = this
    }

    async init() {
        this.browser = await this.getBrowserInfo()
    }

    async getBrowserInfo() {
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

new User()
document.addEventListener('DOMContentLoaded', async () => {
    await self.user.init();
}, false);
