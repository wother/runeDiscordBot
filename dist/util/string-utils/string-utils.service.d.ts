export type StringNumber = 'one' | '1' | 'two' | '2' | 'three' | '3' | 'four' | '4' | 'five' | '5' | null;
export declare class StringUtilsService {
    private BRACKET_REGEX;
    convertStringToNumber(str: StringNumber): number | null;
    private hasBrackets;
    private removeBrackets;
}
