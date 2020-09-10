//author:天灾巫师
//update:2020-3-23

const fs = require('fs');
const path = require('path');

function FileWebpackPlugin(options) {
    this.options = options;
}

FileWebpackPlugin.prototype.apply = function(compiler) {

    //error params
    if (!((this.options) && (this.options.fileList))) {
        console.log("FileWebpackPlugin Error", "can not find fileList object");
        return;
    }

    //get option
    const options = this.options;

    //log
    const Log = function(content) {

        options.log && (console.log('\033[36m' + "[FileWebpackPlugin] \033[39m" + '\033[32m' + content + '\033[39m'));
    }

    //file
    const fileList = options.fileList;
    compiler.hooks.done.tap('FileWebpackPlugin', function(params) {

        //get output dir and hash
        const outputPath = params.compilation.outputOptions.path;
        const hash = params.hash;

        //show result
        //copy file success or fail
        const fileResult = function(name, dist, err) {
            const showContent = err ? err : name + " --> " + dist;
            Log(showContent);
        }

        for (const [fileName, addressList] of Object.entries(fileList)) {
            if (!Array.isArray(addressList)) {
                Log(fileName + " error file address");
                continue;
            }
            let finalName = fileName;

            //file name include hash value
            const hashReg = new RegExp(/\[hash\]|\[hash:{1}\d+\]/);
            const result = finalName.match(hashReg);

            if (result) {
                //find hash
                const hashStr = result[0];
                //max hash length
                const hashCount = Number(hashStr.match(/\d+/) ? hashStr.match(/\d+/)[0] : hash.length)
                finalName = finalName.replace(hashReg, hash.slice(0, hashCount));
            }


            //copy file to all dist
            for (const targetPath of addressList) {
                const fileFrom = path.resolve(outputPath, finalName);
                const fileTo = path.resolve(targetPath, finalName);

                fs.copyFile(fileFrom, fileTo, fileResult.bind(this, finalName, fileTo));

            }

        }

    })

};

module.exports = FileWebpackPlugin;