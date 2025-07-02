export declare class RuneService {
    private readonly logger;
    private FUTHARK_NAMES_ARRAY;
    getFutharkArray(): string[];
    getRandomRune(count: number): {
        name: string;
        imgUrl: string;
        descUrl: string;
    }[];
    runeInfo(name: string): {
        name: string;
        imgUrl: string;
        descUrl: string;
    };
    getRunesInfo(names: string[]): {
        name: string;
        imgUrl: string;
        descUrl: string;
    }[];
    runeInfoImage(name: string): {
        name: string;
        imgUrl: string;
        descUrl: string;
    };
    isRuneName(name: string): boolean;
    private genRuneObject;
    private getImgLink;
    private getInfoImage;
    private genInfoLink;
    private numUniqueRunes;
}
