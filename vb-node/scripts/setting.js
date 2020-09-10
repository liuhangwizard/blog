let config = {
    devMode: false,
    IP: "127.0.0.1",
    port: {
        client: 5000,
        server: 8000,
        rpc: 7000
    },
    directory: {
        //dev
        //client: "./build",
        //prod
        client: "/home/node/vb-blog/public",
        server: "./dist",
        static: "./dist"
    },
}
module.exports = config;