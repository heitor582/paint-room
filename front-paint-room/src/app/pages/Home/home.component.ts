import { HttpResponse } from '@angular/common/http';
import { Component, OnDestroy, OnInit } from '@angular/core';
import { IBodyDtoInput } from 'src/app/shared/interfaces/BodyDtoInput.interface';
import { ICansInfo } from 'src/app/shared/interfaces/InfoCans.interface';
import { HomeService } from 'src/app/shared/services/home.service';

import { takeUntilDestroy } from './../../shared/utils/take-until-destroy';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent implements OnInit, OnDestroy {
  info: HttpResponse<ICansInfo>;
  walls: IBodyDtoInput[] = [];
  numberOfWalls: number = 4;
  submitEnable: boolean = false;
  errorMessage: string = "";
  constructor(private homeService: HomeService) {}

  ngOnInit(): void {}

  public nextWall(event: IBodyDtoInput, index: number): void {
    this.walls[index] = event;
    if (index + 1 == this.numberOfWalls) {
      this.submitEnable = true;
    }
  }

  public getInfoAboutHowManyCans(): void {
    this.homeService
      .getInfoAboutHowManyCans(this.walls)
      .pipe(takeUntilDestroy(this))
      .subscribe(
        (res) => {
          this.info = res;
          this.errorMessage = "";
        },
        (error) => {
          this.errorMessage = error.error
        }
      );
  }

  ngOnDestroy(): void {}
}
