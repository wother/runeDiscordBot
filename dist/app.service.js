"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
var AppService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppService = void 0;
const common_1 = require("@nestjs/common");
const necord_1 = require("necord");
const discord_js_1 = require("discord.js");
const rune_service_1 = require("./rune/rune.service");
const string_utils_service_1 = require("./util/string-utils/string-utils.service");
let AppService = AppService_1 = class AppService {
    constructor(runeService, stringUtils) {
        this.runeService = runeService;
        this.stringUtils = stringUtils;
        this.logger = new common_1.Logger(AppService_1.name);
        this.helpMessage = `The Rune Secrets Bot will draw runes for you from the Elder Futhark.\n
            **Commands are:** \n
            \`!help\` for this help text\n
            \`!cast\` or \`!cast one\` for a single rune casting\n
            \`!cast three\` for a three rune casting\n
            \`!cast five\` for a five rune casting (careful...)\n
            \`!info allrunes\` or \`names\` or \`all\` or \`list\` for a list of all the rune names.\n
            \`!info [runeName]\` for information on a specific Rune.\n
            "draw" "rune" and "cast" are synonymous, eg: \`!draw\` is the same as \`!cast\` or \`!rune.\`
            `;
        this.allRunesCommands = ['allrunes', 'names', 'all', 'list'];
    }
    onReady([client]) {
        this.logger.log(`Bot logged in as ${client.user.username}`);
    }
    onWarn([message]) {
        this.logger.warn(message);
    }
    onPing([message], args) {
        this.logger.log('Received a message');
        return message.reply('pong!');
    }
    onHelp([message], args) {
        this.logger.log('Received a help');
        return message.reply(this.helpMessage);
    }
    onCastRune([message], args) {
        return this.runeCasting(message, args);
    }
    onDrawRune([message], args) {
        return this.runeCasting(message, args);
    }
    onRune([message], args) {
        return this.runeCasting(message, args);
    }
    async onInfo([message], args) {
        this.logger.log('Received a info rune');
        if (args.length === 0) {
            return message.reply(this.helpMessage);
        }
        const runeName = args[0].toLowerCase();
        if (this.allRunesCommands.includes(runeName)) {
            const runes = this.runeService.getRunesInfo(this.runeService.getFutharkArray());
            const runeEmbeds = this.runeToEmbeded(runes);
            let bundle = [];
            for (const rune of runeEmbeds) {
                const index = runeEmbeds.indexOf(rune);
                bundle.push(rune);
                if ((index !== 0 && index % 9 === 0) ||
                    index === runeEmbeds.length - 1) {
                    await message.channel.send({ embeds: bundle });
                    bundle = [];
                }
            }
            return;
        }
        const rune = this.runeService.runeInfo(runeName);
        if (!rune) {
            return message.reply(`The rune ${runeName} is not recognized.`);
        }
        const runeEmbeds = this.runeToEmbeded([rune]);
        return message.channel.send({ embeds: runeEmbeds });
    }
    runeCasting(message, args) {
        var _a;
        this.logger.log('Received a casting rune');
        const count = (_a = this.stringUtils.convertStringToNumber(args[0].toLowerCase())) !== null && _a !== void 0 ? _a : 1;
        const runes = this.runeService.getRandomRune(count);
        const runeEmbeds = this.runeToEmbeded(runes);
        return message.channel.send({ embeds: runeEmbeds });
    }
    runeToEmbeded(runes) {
        return runes.map((rune) => {
            return new discord_js_1.EmbedBuilder()
                .setTitle(rune.name)
                .setURL(rune.descUrl)
                .setImage(rune.imgUrl)
                .setDescription(`You drew ${rune.name}`);
        });
    }
};
__decorate([
    (0, necord_1.Once)('ready'),
    __param(0, (0, necord_1.Context)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onReady", null);
__decorate([
    (0, necord_1.On)('warn'),
    __param(0, (0, necord_1.Context)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onWarn", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'ping',
        description: 'Ping command',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onPing", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'help',
        description: 'This help text',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onHelp", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'cast',
        description: 'for a rune casting',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onCastRune", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'draw',
        description: 'for a rune casting',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onDrawRune", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'rune',
        description: 'for a rune casting',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", void 0)
], AppService.prototype, "onRune", null);
__decorate([
    (0, necord_1.TextCommand)({
        name: 'info',
        description: 'for information on specific Rune or list all runes',
    }),
    __param(0, (0, necord_1.Context)()),
    __param(1, (0, necord_1.Arguments)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Array, Array]),
    __metadata("design:returntype", Promise)
], AppService.prototype, "onInfo", null);
AppService = AppService_1 = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [rune_service_1.RuneService,
        string_utils_service_1.StringUtilsService])
], AppService);
exports.AppService = AppService;
//# sourceMappingURL=app.service.js.map