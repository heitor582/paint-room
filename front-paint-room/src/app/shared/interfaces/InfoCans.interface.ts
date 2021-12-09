export interface ICansInfo {
  cans:                 Can[];
  totalPaitfulRoomArea: number;
  totalLitersOfCans:    number;
}

export interface Can {
  label:    string;
  value:    number;
  quantity: number;
}
