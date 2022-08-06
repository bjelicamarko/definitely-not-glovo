export interface MapAddress {
    display_name: string;
    address: {
        city: string,
        country: string,
        house_number: string,
        postcode: string,
        road: string
    };
}