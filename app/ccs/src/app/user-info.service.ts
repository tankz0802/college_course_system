import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserInfoService {

  private role:string;
  private name:string;
  private id:string;
  constructor() { }

  setRole(role:string) {
    this.role = role;
  }

  getRole() {
    return this.role;
  }

  setName(name:string) {
    this.name = name;
  }

  getName() {
    return this.name;
  }

  setId(id:string) {
    this.id = id;
  }

  getId() {
    return this.id;
  }
}
