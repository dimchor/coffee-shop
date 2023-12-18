import { EventEmitter, Injectable, Output } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class FilteredListService {

  @Output() aClickedEvent = new EventEmitter<string>();

  AClicked(msg: string) {
    this.aClickedEvent.emit(msg);
  }
}
