/* Do not change, this code is generated from Golang structs */

export {};

export class question {
    key?: number;
    value?: string;
    answer?: string;

    static createFrom(source: any = {}) {
        return new question(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.key = source["key"];
        this.value = source["value"];
        this.answer = source["answer"];
    }
}
