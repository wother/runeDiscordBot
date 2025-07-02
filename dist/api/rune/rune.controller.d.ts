import { RuneService } from '../../rune/rune.service';
export declare class RuneController {
    private readonly runeService;
    constructor(runeService: RuneService);
    private readonly logger;
    getRandomRunes(count: number): Promise<{
        name: string;
        imgUrl: string;
        descUrl: string;
    }[]>;
    getRandomRune(name: string): Promise<boolean>;
    runeInfo(name: string): Promise<{
        name: string;
        imgUrl: string;
        descUrl: string;
    }>;
    getFutharkArray(): Promise<string[]>;
}
