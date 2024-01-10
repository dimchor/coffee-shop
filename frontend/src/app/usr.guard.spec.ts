import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { usrGuard } from './usr.guard';

describe('usrGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => usrGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
