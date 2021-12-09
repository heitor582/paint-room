import { Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IBodyDtoInput } from '../../interfaces/BodyDtoInput.interface';

@Component({
  selector: 'app-card',
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.css'],
})
export class CardComponent implements OnInit {
  @Input() index: number;
  @Output() next = new EventEmitter<IBodyDtoInput>();

  form: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.createForm();
  }

  public nextCard() {
    if(this.form.enabled){
      this.form.disable();
      this.next.emit(this.form.getRawValue());
    }else{
      this.form.enable()
    }
  }

  public returnsNumber(name) {
    const value = this.form.controls[name].value;
    if(!isNaN(+value)){
      this.form.controls[name].setValue(+value)
    }
  }

  private createForm(): void {
    this.form = this.fb.group({
      doorsNumber: [0, [Validators.required, Validators.pattern("^[0-9]*$")]],
      windownsNumber: [0, [Validators.required, Validators.pattern("^[0-9]*$")]],
      height: [
        0,
        [Validators.required, Validators.min(1), Validators.pattern("^[0-9]*$")],
      ],
      width: [
        0,
        [Validators.required, Validators.min(1), Validators.pattern("^[0-9]*$")],
      ],
    });
  }
}
