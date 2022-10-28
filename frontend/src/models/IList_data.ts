import { EquipmentInterface } from "./IEquipment";
import { Borrow_cardInterface} from "./IBorrow_card";
import { RoomInterface } from "./IRoom";

export interface List_dataInterface {
  ID?: number;
  Borrow_time: Date | null;
  Return_time: Date | null;

  Equipment_id?: number;
 Equipment?: EquipmentInterface;
  Borrow_card_id?: number;
 Borrow_card?: Borrow_cardInterface;
  Room_id?: number;
  room?: RoomInterface;
}