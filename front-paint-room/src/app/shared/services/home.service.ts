import { HttpClient, HttpEvent, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { APP_CONFIG } from '../../../assets/app.config';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  constructor(public _http: HttpClient) {}

  public getInfoAboutHowManyCans(body: any): Observable<HttpResponse<any>> {
    const url = `${APP_CONFIG.APP_SERVER}/room`;
    return this._http.post<HttpResponse<any>>(url, { walls: body });
  }
}
