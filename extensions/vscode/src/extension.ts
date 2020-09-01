/*
	@Authror: Ayat Maulana
	@Email: me@ayatmaulana.com
*/

import * as vscode from 'vscode';
import { lookpath } from 'lookpath';
import { promisify } from 'util';
import { readFile } from 'fs';
import getos = require('getos');

const getOsPromise = promisify(getos);
const readFilePromise = promisify(readFile);

async function executeTerminal (terminalName: string, command: string) : Promise<vscode.Terminal> {
		const terminal = vscode.window.createTerminal(terminalName);
		terminal.sendText(command);
		terminal.show();
		return terminal;
}

async function checkTadaRunnerGeneratorCore(): Promise<boolean> {
	const check = await lookpath('tada-runner-generator');
	return check ? true : false;
}

async function installTadaRunnerGeneratorCore(): Promise<string|undefined> {
	const isCoreExist: boolean = await checkTadaRunnerGeneratorCore();
	if (!isCoreExist) {
		throw new Error('Please Install core first, more info https://github.com/ayatmaulana/tada-runner-generator');
	}

	const detail: getos.Os = await getOsPromise();
	switch(detail.os) {
		case 'darwin':
			const command: string = 'brew tap ayatmaulana/pkg && brew install tada-runner-generator';
			await executeTerminal('Installing Tada Runner Generator Core', command);
			return "✅ Success installing tada-runner-generator"; 

		case 'linux':
			const linuxOs: getos.LinuxOs = (detail as getos.LinuxOs);
			const { dist } = linuxOs;

			if (['arch', 'manjaro']) {
				const command: string = 'yay -S tada-runner-generator';
				await executeTerminal('Installing Tada Runner Generator Core', command);
				return "✅ Success installing tada-runner-generator";
			} else {
				throw new Error(`
					You Are using Linux with distro ${dist},
					currently we did not support with your distro
					please install manually from go source or download binary from

					github.com/ayatmaulana/tada-runner-generator/releases
				`);
			}
	}
}

async function checkIfInProjectFolder(): Promise<boolean> {
	const rootPath = vscode.workspace.rootPath;
	const readJson = await readFilePromise(rootPath + '/package.json',  'utf8');
	const packageJson = JSON.parse(readJson);

	return (packageJson.name === 'runners');
}


// this method is called when your extension is activated
// your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {

	console.log('Congratulations, your extension "tada-runner-generator" is now active!');

	const checkCore: vscode.Disposable = vscode.commands.registerCommand('tada-runner-generator.checkCore', async () => {
		const check = await checkTadaRunnerGeneratorCore();
		if(check) {
			vscode.window.showInformationMessage("Core Already Installed on your machine !!");
		} else {
			vscode.window.showInformationMessage("Core is not installed  on your machine, please install first before use this extension");
		}
	});

	const installCore: vscode.Disposable = vscode.commands.registerCommand('tada-runner-generator.installCore', async () => {
		try {
			const installCoreProcess: string = await installTadaRunnerGeneratorCore() as string;
			vscode.window.showInformationMessage(installCoreProcess);
		} catch(error) {
			vscode.window.showInformationMessage(error.message);
		}
	});

	const addNewRunner: vscode.Disposable = vscode.commands.registerCommand('tada-runner-generator.addNewRunner', async () => {
		try {
			const checkFolder: boolean = await checkIfInProjectFolder();
			console.log('checkFolder: ', checkFolder);
			if (!checkFolder) {
				throw new Error('Please run into your runner project !');
			}
	
			const name = await vscode.window.showInputBox({
				placeHolder: 'Enter runner name',
			});
			if (typeof name === 'undefined') {
				return;
			}
	
			const isCopyDir = await vscode.window.showQuickPick(['Yes', 'No'], {
				placeHolder: 'Copy dir app,config,locale,lib to your runner dir ?',
			});
			if (typeof isCopyDir === 'undefined') {
				return;
			}
	
			const npmInstall = await vscode.window.showQuickPick(['Yes', 'No'], {
				placeHolder: 'Do npm install ?',
			});
			if (typeof npmInstall === 'undefined') {
				return;
			}
	
			let additionalCondition : string = isCopyDir === 'Yes' ? '-c ' : '';
			additionalCondition = npmInstall === 'Yes' ? additionalCondition+'-i' : additionalCondition;
			await executeTerminal("Create Runner", `tada-runner-generator add '${name}'  ${additionalCondition}`);
	
		} catch(error) {
			console.log('error: ', error.message);
			vscode.window.showInformationMessage(error.message);
		}
	});
	
	context.subscriptions.push(checkCore);
	context.subscriptions.push(installCore);
	context.subscriptions.push(addNewRunner);
}

// this method is called when your extension is deactivated
export function deactivate() {}
