/**
 *  Copyright 2015 Grafana Labs
 *  Modifications Copyright 2019 Tornimo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {TornimoDatasource} from './datasource';
import {TornimoQueryCtrl} from './query_ctrl';
import {TornimoConfigCtrl} from './config_ctrl';

class TornimoAnnotationsQueryCtrl {
    static templateUrl = 'partials/annotations.editor.html';
}

export {
    TornimoDatasource as Datasource,
    TornimoQueryCtrl as QueryCtrl,
    TornimoConfigCtrl as ConfigCtrl,
    TornimoAnnotationsQueryCtrl as AnnotationsQueryCtrl,
};
