"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var RuneService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.RuneService = void 0;
const common_1 = require("@nestjs/common");
let RuneService = RuneService_1 = class RuneService {
    constructor() {
        this.logger = new common_1.Logger(RuneService_1.name);
        this.FUTHARK_NAMES_ARRAY = [
            'algiz',
            'ansuz',
            'berkano',
            'dagaz',
            'ehwaz',
            'fehu',
            'gebo',
            'hagalaz',
            'ihwaz',
            'inguz',
            'isa',
            'jera',
            'kenaz',
            'laguz',
            'mannaz',
            'nauthiz',
            'othala',
            'perthro',
            'raidho',
            'sowilo',
            'thurisaz',
            'tiwaz',
            'uruz',
            'wunjo',
        ];
    }
    getFutharkArray() {
        return this.FUTHARK_NAMES_ARRAY;
    }
    getRandomRune(count) {
        this.logger.log(`getting ${count} runes`);
        if (1 < count && count <= 5) {
            return this.numUniqueRunes(count);
        }
        return this.numUniqueRunes(1);
    }
    runeInfo(name) {
        this.logger.log(`getting rune info for ${name}`);
        if (this.isRuneName(name))
            return this.runeInfoImage(name);
        return null;
    }
    getRunesInfo(names) {
        return names
            .map((name) => {
            if (this.isRuneName(name)) {
                this.logger.log(`rune: ${name}`);
                return this.runeInfoImage(name);
            }
            return null;
        })
            .filter((rune) => rune);
    }
    runeInfoImage(name) {
        this.logger.log(`getting rune info image for ${name}`);
        return {
            name: name,
            imgUrl: this.getInfoImage(name),
            descUrl: this.genInfoLink(name),
        };
    }
    isRuneName(name) {
        this.logger.log(`checking if ${name} is a rune`);
        return this.FUTHARK_NAMES_ARRAY.includes(name);
    }
    genRuneObject(name) {
        this.logger.log(`generating rune object for ${name}`);
        return {
            name: name,
            imgUrl: this.getImgLink(name),
            descUrl: this.genInfoLink(name),
        };
    }
    getImgLink(name) {
        return `https://runesecrets.com/img/${name}-100x100.gif`;
    }
    getInfoImage(name) {
        return `https://runesecrets.com/img/info-${name}.gif`;
    }
    genInfoLink(name) {
        return `https://runesecrets.com/rune-meanings/${name}`;
    }
    numUniqueRunes(count) {
        let futharkArrayCopy = [...this.FUTHARK_NAMES_ARRAY];
        let randoRunes = futharkArrayCopy
            .sort(() => {
            return 0.5 - Math.random();
        })
            .slice(0, count);
        return randoRunes.map((runeName) => this.genRuneObject(runeName));
    }
};
RuneService = RuneService_1 = __decorate([
    (0, common_1.Injectable)()
], RuneService);
exports.RuneService = RuneService;
//# sourceMappingURL=rune.service.js.map