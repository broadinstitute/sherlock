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
/**
 * 
 * @export
 * @interface V2controllersCreatableChartRelease
 */
export interface V2controllersCreatableChartRelease {
    /**
     * When creating, will default to the app's mainline branch if no other app version info is present
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    appVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    appVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    appVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    appVersionFollowChartRelease?: string;
    /**
     * // When creating, will default to automatically reference any provided app version fields
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    appVersionResolver?: V2controllersCreatableChartReleaseAppVersionResolverEnum;
    /**
     * Required when creating
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    chart?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    chartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    chartVersionFollowChartRelease?: string;
    /**
     * When creating, will default to automatically reference any provided chart version
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    chartVersionResolver?: V2controllersCreatableChartReleaseChartVersionResolverEnum;
    /**
     * When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    cluster?: string;
    /**
     * Either this or cluster must be provided.
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    environment?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    firecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    helmfileRef?: string;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    name?: string;
    /**
     * When creating, will default to the environment's default namespace, if provided
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    namespace?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    pagerdutyIntegration?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {number}
     * @memberof V2controllersCreatableChartRelease
     */
    port?: number;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    protocol?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof V2controllersCreatableChartRelease
     */
    subdomain?: string;
}


/**
 * @export
 */
export const V2controllersCreatableChartReleaseAppVersionResolverEnum = {
    Branch: 'branch',
    Commit: 'commit',
    Exact: 'exact',
    Follow: 'follow',
    None: 'none'
} as const;
export type V2controllersCreatableChartReleaseAppVersionResolverEnum = typeof V2controllersCreatableChartReleaseAppVersionResolverEnum[keyof typeof V2controllersCreatableChartReleaseAppVersionResolverEnum];

/**
 * @export
 */
export const V2controllersCreatableChartReleaseChartVersionResolverEnum = {
    Latest: 'latest',
    Exact: 'exact',
    Follow: 'follow'
} as const;
export type V2controllersCreatableChartReleaseChartVersionResolverEnum = typeof V2controllersCreatableChartReleaseChartVersionResolverEnum[keyof typeof V2controllersCreatableChartReleaseChartVersionResolverEnum];


/**
 * Check if a given object implements the V2controllersCreatableChartRelease interface.
 */
export function instanceOfV2controllersCreatableChartRelease(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersCreatableChartReleaseFromJSON(json: any): V2controllersCreatableChartRelease {
    return V2controllersCreatableChartReleaseFromJSONTyped(json, false);
}

export function V2controllersCreatableChartReleaseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersCreatableChartRelease {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appVersionBranch': !exists(json, 'appVersionBranch') ? undefined : json['appVersionBranch'],
        'appVersionCommit': !exists(json, 'appVersionCommit') ? undefined : json['appVersionCommit'],
        'appVersionExact': !exists(json, 'appVersionExact') ? undefined : json['appVersionExact'],
        'appVersionFollowChartRelease': !exists(json, 'appVersionFollowChartRelease') ? undefined : json['appVersionFollowChartRelease'],
        'appVersionResolver': !exists(json, 'appVersionResolver') ? undefined : json['appVersionResolver'],
        'chart': !exists(json, 'chart') ? undefined : json['chart'],
        'chartVersionExact': !exists(json, 'chartVersionExact') ? undefined : json['chartVersionExact'],
        'chartVersionFollowChartRelease': !exists(json, 'chartVersionFollowChartRelease') ? undefined : json['chartVersionFollowChartRelease'],
        'chartVersionResolver': !exists(json, 'chartVersionResolver') ? undefined : json['chartVersionResolver'],
        'cluster': !exists(json, 'cluster') ? undefined : json['cluster'],
        'environment': !exists(json, 'environment') ? undefined : json['environment'],
        'firecloudDevelopRef': !exists(json, 'firecloudDevelopRef') ? undefined : json['firecloudDevelopRef'],
        'helmfileRef': !exists(json, 'helmfileRef') ? undefined : json['helmfileRef'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
        'pagerdutyIntegration': !exists(json, 'pagerdutyIntegration') ? undefined : json['pagerdutyIntegration'],
        'port': !exists(json, 'port') ? undefined : json['port'],
        'protocol': !exists(json, 'protocol') ? undefined : json['protocol'],
        'subdomain': !exists(json, 'subdomain') ? undefined : json['subdomain'],
    };
}

export function V2controllersCreatableChartReleaseToJSON(value?: V2controllersCreatableChartRelease | null): any {
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
        'appVersionFollowChartRelease': value.appVersionFollowChartRelease,
        'appVersionResolver': value.appVersionResolver,
        'chart': value.chart,
        'chartVersionExact': value.chartVersionExact,
        'chartVersionFollowChartRelease': value.chartVersionFollowChartRelease,
        'chartVersionResolver': value.chartVersionResolver,
        'cluster': value.cluster,
        'environment': value.environment,
        'firecloudDevelopRef': value.firecloudDevelopRef,
        'helmfileRef': value.helmfileRef,
        'name': value.name,
        'namespace': value.namespace,
        'pagerdutyIntegration': value.pagerdutyIntegration,
        'port': value.port,
        'protocol': value.protocol,
        'subdomain': value.subdomain,
    };
}

