import { NgModule } from "@angular/core";
import { CommonModule } from '@angular/common';

import {MatSnackBarModule} from '@angular/material/snack-bar';
import { SnackBarService } from "./services/snack-bar.service";
import { MaterialExampleModule } from 'src/material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
    declarations: [

    ],
    imports: [
        CommonModule,
        MaterialExampleModule,
        MatSnackBarModule,
        FormsModule,
        ReactiveFormsModule
    ],
    exports: [

    ],
    providers: [
        SnackBarService,
    ]
})
export class SharedModule { }