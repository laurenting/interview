import * as models from './models';

export interface go {
  "main": {
    "App": {
		Greet(arg1:string):Promise<string>
		OpenQuestionnaire():Promise<Array<models.question>|Error>
		Save(arg1:Array<models.question>):Promise<string|Error>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
