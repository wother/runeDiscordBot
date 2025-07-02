import { ContextOf, TextCommandContext } from 'necord';
import { RuneService } from './rune/rune.service';
import { StringUtilsService } from './util/string-utils/string-utils.service';
export declare class AppService {
    private readonly runeService;
    private readonly stringUtils;
    constructor(runeService: RuneService, stringUtils: StringUtilsService);
    private readonly logger;
    private readonly helpMessage;
    private readonly allRunesCommands;
    onReady([client]: ContextOf<'ready'>): void;
    onWarn([message]: ContextOf<'warn'>): void;
    onPing([message]: TextCommandContext, args: string[]): Promise<import("discord.js").Message<boolean>>;
    onHelp([message]: TextCommandContext, args: string[]): Promise<import("discord.js").Message<boolean>>;
    onCastRune([message]: TextCommandContext, args: string[]): any;
    onDrawRune([message]: TextCommandContext, args: string[]): any;
    onRune([message]: TextCommandContext, args: string[]): any;
    onInfo([message]: TextCommandContext, args: string[]): Promise<import("discord.js").Message<boolean>>;
    private runeCasting;
    private runeToEmbeded;
}
