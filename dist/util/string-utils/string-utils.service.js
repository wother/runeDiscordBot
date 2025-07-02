"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.StringUtilsService = void 0;
const common_1 = require("@nestjs/common");
let StringUtilsService = class StringUtilsService {
    constructor() {
        this.BRACKET_REGEX = /([\[\]])\w*/g;
    }
    convertStringToNumber(str) {
        let count = this.hasBrackets(str) ? this.removeBrackets(str) : str;
        switch (count) {
            case '1':
            case 'one':
                return 1;
            case '2':
            case 'two':
                return 2;
            case '3':
            case 'three':
                return 3;
            case '4':
            case 'four':
                return 4;
            case '5':
            case 'five':
                return 5;
            default:
                return null;
        }
    }
    hasBrackets(str) {
        return this.BRACKET_REGEX.test(str);
    }
    removeBrackets(str) {
        let output = '';
        let strArray = str.split('');
        strArray.forEach((stringChar) => {
            if (!this.BRACKET_REGEX.test(stringChar)) {
                output += stringChar;
            }
        });
        return output;
    }
};
StringUtilsService = __decorate([
    (0, common_1.Injectable)()
], StringUtilsService);
exports.StringUtilsService = StringUtilsService;
//# sourceMappingURL=string-utils.service.js.map