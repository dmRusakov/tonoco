class User {
    constructor() {
        self.user = this
    }

    init() {
        this.location = async () => {
            try {
                console.log(222);
                const response = await fetch("http://ip-api.com/json/");
                console.log(response);
                const data =  await response.json();

                console.log(data);


                return {

                }
            } catch(error) {
                console.error("Error fetching user data:", error);
                return null;
            }
        }
    }
}

new User()
document.addEventListener('DOMContentLoaded', async () => {
    self.user.init();
}, false);
