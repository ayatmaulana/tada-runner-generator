"use strict";
/*
    @Authror: Ayat Maulana
    @Email: me@ayatmaulana.com
*/
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.deactivate = exports.activate = void 0;
const vscode = require("vscode");
const lookpath_1 = require("lookpath");
const util_1 = require("util");
const fs_1 = require("fs");
const getos = require("getos");
const getOsPromise = util_1.promisify(getos);
const readFilePromise = util_1.promisify(fs_1.readFile);
function executeTerminal(terminalName, command) {
    return __awaiter(this, void 0, void 0, function* () {
        const terminal = vscode.window.createTerminal(terminalName);
        terminal.sendText(command);
        terminal.show();
        return terminal;
    });
}
function checkTadaRunnerGeneratorCore() {
    return __awaiter(this, void 0, void 0, function* () {
        const check = yield lookpath_1.lookpath('tada-runner-generator');
        return check ? true : false;
    });
}
function installTadaRunnerGeneratorCore() {
    return __awaiter(this, void 0, void 0, function* () {
        const isCoreExist = yield checkTadaRunnerGeneratorCore();
        if (!isCoreExist) {
            throw new Error('Please Install core first, more info https://github.com/ayatmaulana/tada-runner-generator');
        }
        const detail = yield getOsPromise();
        switch (detail.os) {
            case 'darwin':
                const command = 'brew tap ayatmaulana/pkg && brew install tada-runner-generator';
                yield executeTerminal('Installing Tada Runner Generator Core', command);
                return "✅ Success installing tada-runner-generator";
            case 'linux':
                const linuxOs = detail;
                const { dist } = linuxOs;
                if (['arch', 'manjaro']) {
                    const command = 'yay -S tada-runner-generator';
                    yield executeTerminal('Installing Tada Runner Generator Core', command);
                    return "✅ Success installing tada-runner-generator";
                }
                else {
                    throw new Error(`
					You Are using Linux with distro ${dist},
					currently we did not support with your distro
					please install manually from go source or download binary from

					github.com/ayatmaulana/tada-runner-generator/releases
				`);
                }
        }
    });
}
function checkIfInProjectFolder() {
    return __awaiter(this, void 0, void 0, function* () {
        const rootPath = vscode.workspace.rootPath;
        const readJson = yield readFilePromise(rootPath + '/package.json', 'utf8');
        const packageJson = JSON.parse(readJson);
        return (packageJson.name === 'runners');
    });
}
// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
function activate(context) {
    console.log('Congratulations, your extension "tada-runner-generator" is now active!');
    const checkCore = vscode.commands.registerCommand('tada-runner-generator.checkCore', () => __awaiter(this, void 0, void 0, function* () {
        const check = yield checkTadaRunnerGeneratorCore();
        if (check) {
            vscode.window.showInformationMessage("Core Already Installed on your machine !!");
        }
        else {
            vscode.window.showInformationMessage("Core is not installed  on your machine, please install first before use this extension");
        }
    }));
    const installCore = vscode.commands.registerCommand('tada-runner-generator.installCore', () => __awaiter(this, void 0, void 0, function* () {
        try {
            const installCoreProcess = yield installTadaRunnerGeneratorCore();
            vscode.window.showInformationMessage(installCoreProcess);
        }
        catch (error) {
            vscode.window.showInformationMessage(error.message);
        }
    }));
    const addNewRunner = vscode.commands.registerCommand('tada-runner-generator.addNewRunner', () => __awaiter(this, void 0, void 0, function* () {
        try {
            const checkFolder = yield checkIfInProjectFolder();
            console.log('checkFolder: ', checkFolder);
            if (!checkFolder) {
                throw new Error('Please run into your runner project !');
            }
            const name = yield vscode.window.showInputBox({
                placeHolder: 'Enter runner name',
            });
            if (typeof name === 'undefined') {
                return;
            }
            const isCopyDir = yield vscode.window.showQuickPick(['Yes', 'No'], {
                placeHolder: 'Copy dir app,config,locale,lib to your runner dir ?',
            });
            if (typeof isCopyDir === 'undefined') {
                return;
            }
            const npmInstall = yield vscode.window.showQuickPick(['Yes', 'No'], {
                placeHolder: 'Do npm install ?',
            });
            if (typeof npmInstall === 'undefined') {
                return;
            }
            let additionalCondition = isCopyDir === 'Yes' ? '-c ' : '';
            additionalCondition = npmInstall === 'Yes' ? additionalCondition + '-i' : additionalCondition;
            yield executeTerminal("Create Runner", `tada-runner-generator add '${name}'  ${additionalCondition}`);
        }
        catch (error) {
            console.log('error: ', error.message);
            vscode.window.showInformationMessage(error.message);
        }
    }));
    context.subscriptions.push(checkCore);
    context.subscriptions.push(installCore);
    context.subscriptions.push(addNewRunner);
}
exports.activate = activate;
// this method is called when your extension is deactivated
function deactivate() { }
exports.deactivate = deactivate;
//# sourceMappingURL=extension.js.map