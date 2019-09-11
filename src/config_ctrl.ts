export class TornimoConfigCtrl {
    static templateUrl = 'partials/config.html';
    datasourceSrv: any;
    current: any;

    /** @ngInject */
    constructor($scope, datasourceSrv) {
        this.datasourceSrv = datasourceSrv;
        this.current.jsonData = this.current.jsonData || {};
        this.current.jsonData.tornimoVersion = this.current.jsonData.tornimoVersion || '1.0';
        this.current.jsonData.bucketSizeSeconds = this.current.jsonData.bucketSizeSeconds || 60;

        this.autoDetectTornimoVersion();
    }

    autoDetectTornimoVersion() {
        if (!this.current.id) {
            return;
        }

        this.datasourceSrv
            .loadDatasource(this.current.name)
            .then(ds => {
                return ds.getVersion();
            })
            .then(version => {
                this.tornimoVersions.push({ name: version, value: version });
                this.current.jsonData.tornimoVersion = version;
            });
    }

    tornimoVersions = [{ name: '0.382', value: '0.382' }];

    bucketSizes = [
        { name: '1s', value: '1' },
        { name: '5s', value: '5' },
        { name: '10s', value: '10' },
        { name: '15s', value: '15' },
        { name: '30s', value: '30' },
        { name: '60s', value: '60' },
    ];
}
