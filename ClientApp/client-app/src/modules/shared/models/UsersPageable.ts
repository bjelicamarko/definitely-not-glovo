import { UserDTO } from "./UserDTO";

export interface UsersPageable {
    Elements: UserDTO[];
    TotalElements: number;
}