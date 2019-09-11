import {TornimoDatasource} from './datasource';
import {TornimoQueryCtrl} from './query_ctrl';
import {TornimoConfigCtrl} from './config_ctrl';

class HyperionAnnotationsQueryCtrl {
    static templateUrl = 'partials/annotations.editor.html';
}

export {
    TornimoDatasource as Datasource,
    TornimoQueryCtrl as QueryCtrl,
    TornimoConfigCtrl as ConfigCtrl,
    HyperionAnnotationsQueryCtrl as AnnotationsQueryCtrl,
};
