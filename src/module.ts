import {HyperionDatasource} from './datasource';
import {HyperionQueryCtrl} from './query_ctrl';
import {HyperionConfigCtrl} from './config_ctrl';

class HyperionAnnotationsQueryCtrl {
    static templateUrl = 'partials/annotations.editor.html';
}

export {
    HyperionDatasource as Datasource,
    HyperionQueryCtrl as QueryCtrl,
    HyperionConfigCtrl as ConfigCtrl,
    HyperionAnnotationsQueryCtrl as AnnotationsQueryCtrl,
};
