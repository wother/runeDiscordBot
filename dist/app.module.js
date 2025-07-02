"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppModule = void 0;
const common_1 = require("@nestjs/common");
const app_controller_1 = require("./app.controller");
const app_service_1 = require("./app.service");
const necord_1 = require("necord");
const discord_js_1 = require("discord.js");
const config_1 = require("@nestjs/config");
const api_module_1 = require("./api/api.module");
const rune_module_1 = require("./rune/rune.module");
const util_module_1 = require("./util/util.module");
let AppModule = class AppModule {
};
AppModule = __decorate([
    (0, common_1.Module)({
        imports: [
            necord_1.NecordModule.forRoot({
                token: process.env.DISCORD_TOKEN,
                prefix: '!',
                intents: [
                    discord_js_1.IntentsBitField.Flags.Guilds,
                    discord_js_1.IntentsBitField.Flags.GuildMessages,
                    discord_js_1.IntentsBitField.Flags.MessageContent,
                    discord_js_1.IntentsBitField.Flags.GuildPresences,
                ],
            }),
            config_1.ConfigModule.forRoot(),
            api_module_1.ApiModule,
            rune_module_1.RuneModule,
            util_module_1.UtilModule,
        ],
        controllers: [app_controller_1.AppController],
        providers: [app_service_1.AppService],
    })
], AppModule);
exports.AppModule = AppModule;
//# sourceMappingURL=app.module.js.map