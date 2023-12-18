import { TestBed } from '@angular/core/testing';

import { FilteredListService } from './filtered-list.service';

describe('FilteredListService', () => {
  let service: FilteredListService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FilteredListService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
