export class HyperionConfigCtrl {
    static templateUrl = 'partials/config.html';
    // static templateUrl = 'public/app/plugins/datasource/hyperion/partials/config.html';
    datasourceSrv: any;
    current: any;

    /** @ngInject */
    constructor($scope, datasourceSrv) {
        this.datasourceSrv = datasourceSrv;
        this.current.jsonData = this.current.jsonData || {};
        this.current.jsonData.hyperionVersion = this.current.jsonData.hyperionVersion || '1.0';
        this.current.jsonData.bucketSizeSeconds = this.current.jsonData.bucketSizeSeconds || 60;

        this.autoDetecthyperionVersion();
    }

    autoDetecthyperionVersion() {
        if (!this.current.id) {
            return;
        }

        this.datasourceSrv
            .loadDatasource(this.current.name)
            .then(ds => {
                return ds.getVersion();
            })
            .then(version => {
                this.hyperionVersions.push({ name: version, value: version });
                this.current.jsonData.hyperionVersion = version;
            });
    }

    hyperionVersions = [{ name: '0.382', value: '0.382' }];

    bucketSizes = [
        { name: '1s', value: '1' },
        { name: '5s', value: '5' },
        { name: '10s', value: '10' },
        { name: '15s', value: '15' },
        { name: '30s', value: '30' },
        { name: '60s', value: '60' },
    ];
}
