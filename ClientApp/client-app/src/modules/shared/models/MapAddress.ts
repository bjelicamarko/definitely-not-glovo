export interface MapAddress {
    display_name: string;
    address: {
        town: string,
        country: string,
        house_number: string,
        postcode: string,
        road: string
    };
}