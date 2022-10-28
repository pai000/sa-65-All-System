import { EmployeeInterface } from "./IEmployee";
import { GendersInterface } from "./IGender";
import { ProvincesInterface } from "./IProvince";
import { ProgramInterface } from "./IProgram";
import { RoleInterface } from "./IRole";

export interface StudentInterface {
  ID?:      number;
  STUDENT_NUMBER?:      number;
  STUDENT_NAME? :       string;
  PERSONAL_ID? :        string;
  Password?:            string;
  GenderID?:            number;
  Gender?:              GendersInterface;
  ProvinceID?:          number;
  Province?:            ProvincesInterface;
  ProgramID?:           number;
  Program?:             ProgramInterface;
  RoleID?:              number;
  Role?:                RoleInterface;
  EmployeeID? :         number;
  Employee? :           EmployeeInterface;
}