/* tslint:disable */
/* eslint-disable */
/**
 * Sherlock
 * The Data Science Platform\'s source-of-truth service. Note: this API will try to load and return associations in responses, so clients won\'t need to make as many requests. This behavior isn\'t recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).
 *
 * The version of the OpenAPI document: development
 * Contact: dsp-devops@broadinstitute.org
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V2controllersAppVersion } from './V2controllersAppVersion';
import {
    V2controllersAppVersionFromJSON,
    V2controllersAppVersionFromJSONTyped,
    V2controllersAppVersionToJSON,
} from './V2controllersAppVersion';
import type { V2controllersChart } from './V2controllersChart';
import {
    V2controllersChartFromJSON,
    V2controllersChartFromJSONTyped,
    V2controllersChartToJSON,
} from './V2controllersChart';
import type { V2controllersChartVersion } from './V2controllersChartVersion';
import {
    V2controllersChartVersionFromJSON,
    V2controllersChartVersionFromJSONTyped,
    V2controllersChartVersionToJSON,
} from './V2controllersChartVersion';
import type { V2controllersCluster } from './V2controllersCluster';
import {
    V2controllersClusterFromJSON,
    V2controllersClusterFromJSONTyped,
    V2controllersClusterToJSON,
} from './V2controllersCluster';
import type { V2controllersEnvironment } from './V2controllersEnvironment';
import {
    V2controllersEnvironmentFromJSON,
    V2controllersEnvironmentFromJSONTyped,
    V2controllersEnvironmentToJSON,
} from './V2controllersEnvironment';

/**
 * 
 * @export
 * @interface V2controllersChartRelease
 */
export interface V2controllersChartRelease {
    /**
     * When creating, will default to the app's mainline branch if no other app version info is present
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    appVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    appVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    appVersionExact?: string;
    /**
     * 
     * @type {V2controllersAppVersion}
     * @memberof V2controllersChartRelease
     */
    appVersionInfo?: V2controllersAppVersion;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    appVersionReference?: string;
    /**
     * // When creating, will default to automatically reference any provided app version fields
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    appVersionResolver?: V2controllersChartReleaseAppVersionResolverEnum;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    chart?: string;
    /**
     * 
     * @type {V2controllersChart}
     * @memberof V2controllersChartRelease
     */
    chartInfo?: V2controllersChart;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    chartVersionExact?: string;
    /**
     * 
     * @type {V2controllersChartVersion}
     * @memberof V2controllersChartRelease
     */
    chartVersionInfo?: V2controllersChartVersion;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    chartVersionReference?: string;
    /**
     * When creating, will default to automatically reference any provided chart version
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    chartVersionResolver?: V2controllersChartReleaseChartVersionResolverEnum;
    /**
     * When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    cluster?: string;
    /**
     * 
     * @type {V2controllersCluster}
     * @memberof V2controllersChartRelease
     */
    clusterInfo?: V2controllersCluster;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    createdAt?: string;
    /**
     * Calculated field
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    destinationType?: string;
    /**
     * Either this or cluster must be provided.
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    environment?: string;
    /**
     * 
     * @type {V2controllersEnvironment}
     * @memberof V2controllersChartRelease
     */
    environmentInfo?: V2controllersEnvironment;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    firecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    helmfileRef?: string;
    /**
     * 
     * @type {number}
     * @memberof V2controllersChartRelease
     */
    id?: number;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    name?: string;
    /**
     * When creating, will default to the environment's default namespace, if provided
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    namespace?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {number}
     * @memberof V2controllersChartRelease
     */
    port?: number;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    protocol?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    subdomain?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChartRelease
     */
    updatedAt?: string;
}


/**
 * @export
 */
export const V2controllersChartReleaseAppVersionResolverEnum = {
    Branch: 'branch',
    Commit: 'commit',
    Exact: 'exact',
    None: 'none'
} as const;
export type V2controllersChartReleaseAppVersionResolverEnum = typeof V2controllersChartReleaseAppVersionResolverEnum[keyof typeof V2controllersChartReleaseAppVersionResolverEnum];

/**
 * @export
 */
export const V2controllersChartReleaseChartVersionResolverEnum = {
    Latest: 'latest',
    Exact: 'exact'
} as const;
export type V2controllersChartReleaseChartVersionResolverEnum = typeof V2controllersChartReleaseChartVersionResolverEnum[keyof typeof V2controllersChartReleaseChartVersionResolverEnum];


/**
 * Check if a given object implements the V2controllersChartRelease interface.
 */
export function instanceOfV2controllersChartRelease(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersChartReleaseFromJSON(json: any): V2controllersChartRelease {
    return V2controllersChartReleaseFromJSONTyped(json, false);
}

export function V2controllersChartReleaseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersChartRelease {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appVersionBranch': !exists(json, 'appVersionBranch') ? undefined : json['appVersionBranch'],
        'appVersionCommit': !exists(json, 'appVersionCommit') ? undefined : json['appVersionCommit'],
        'appVersionExact': !exists(json, 'appVersionExact') ? undefined : json['appVersionExact'],
        'appVersionInfo': !exists(json, 'appVersionInfo') ? undefined : V2controllersAppVersionFromJSON(json['appVersionInfo']),
        'appVersionReference': !exists(json, 'appVersionReference') ? undefined : json['appVersionReference'],
        'appVersionResolver': !exists(json, 'appVersionResolver') ? undefined : json['appVersionResolver'],
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartInfo': !exists(json, 'chartInfo') ? undefined : V2controllersChartFromJSON(json['chartInfo']),
        'chartVersionExact': !exists(json, 'chartVersionExact') ? undefined : json['chartVersionExact'],
        'chartVersionInfo': !exists(json, 'chartVersionInfo') ? undefined : V2controllersChartVersionFromJSON(json['chartVersionInfo']),
        'chartVersionReference': !exists(json, 'chartVersionReference') ? undefined : json['chartVersionReference'],
        'chartVersionResolver': !exists(json, 'chartVersionResolver') ? undefined : json['chartVersionResolver'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'clusterInfo': !exists(json, 'clusterInfo') ? undefined : V2controllersClusterFromJSON(json['clusterInfo']),
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'destinationType': !exists(json, 'destinationType') ? undefined : json['destinationType'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'environmentInfo': !exists(json, 'environmentInfo') ? undefined : V2controllersEnvironmentFromJSON(json['environmentInfo']),
        'firecloudDevelopRef': !exists(json, 'firecloudDevelopRef') ? undefined : json['firecloudDevelopRef'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
        'port': !exists(json, 'port') ? undefined : json['port'],
        'protocol': !exists(json, 'protocol') ? undefined : json['protocol'],
        'subdomain': !exists(json, 'subdomain') ? undefined : json['subdomain'],
        'updatedAt': !exists(json, 'updatedAt') ? undefined : json['updatedAt'],
    };
}

export function V2controllersChartReleaseToJSON(value?: V2controllersChartRelease | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appVersionBranch': value.appVersionBranch,
        'appVersionCommit': value.appVersionCommit,
        'appVersionExact': value.appVersionExact,
        'appVersionInfo': V2controllersAppVersionToJSON(value.appVersionInfo),
        'appVersionReference': value.appVersionReference,
        'appVersionResolver': value.appVersionResolver,
        'chart': value.chart,
        'chartInfo': V2controllersChartToJSON(value.chartInfo),
        'chartVersionExact': value.chartVersionExact,
        'chartVersionInfo': V2controllersChartVersionToJSON(value.chartVersionInfo),
        'chartVersionReference': value.chartVersionReference,
        'chartVersionResolver': value.chartVersionResolver,
        'cluster': value.cluster,
        'clusterInfo': V2controllersClusterToJSON(value.clusterInfo),
        'createdAt': value.createdAt,
        'destinationType': value.destinationType,
        'environment': value.environment,
        'environmentInfo': V2controllersEnvironmentToJSON(value.environmentInfo),
        'firecloudDevelopRef': value.firecloudDevelopRef,
        'helmfileRef': value.helmfileRef,
        'id': value.id,
        'name': value.name,
        'namespace': value.namespace,
        'port': value.port,
        'protocol': value.protocol,
        'subdomain': value.subdomain,
        'updatedAt': value.updatedAt,
    };
}

