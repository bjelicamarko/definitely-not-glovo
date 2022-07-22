import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';

import {MatSnackBarModule} from '@angular/material/snack-bar';
import { SnackBarService } from "./services/snack-bar.service";
import { MaterialExampleModule } from 'src/material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { PaginationComponent } from './components/pagination/pagination.component';
import { UtilsService } from "./services/utils.service";
import { ConformationDialogComponent } from './components/conformation-dialog/conformation-dialog.component';

@NgModule({
    declarations: [

    
    PaginationComponent,
            ConformationDialogComponent
  ],
    imports: [
        CommonModule,
        MaterialExampleModule,
        MatSnackBarModule,
        FormsModule,
        ReactiveFormsModule
    ],
    exports: [
        PaginationComponent
    ],
    providers: [
        SnackBarService,
        UtilsService,
    ]
})
export class SharedModule { }