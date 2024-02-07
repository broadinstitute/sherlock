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
 * @interface SherlockChartReleaseV3Create
 */
export interface SherlockChartReleaseV3Create {
    /**
     * When creating, will default to the app's mainline branch if no other app version info is present
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    appVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    appVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    appVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    appVersionFollowChartRelease?: string;
    /**
     * // When creating, will default to automatically reference any provided app version fields
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    appVersionResolver?: SherlockChartReleaseV3CreateAppVersionResolverEnum;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    chart?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    chartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    chartVersionFollowChartRelease?: string;
    /**
     * When creating, will default to automatically reference any provided chart version
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    chartVersionResolver?: SherlockChartReleaseV3CreateChartVersionResolverEnum;
    /**
     * When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    cluster?: string;
    /**
     * Either this or cluster must be provided.
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    environment?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    firecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    helmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartReleaseV3Create
     */
    helmfileRefEnabled?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartReleaseV3Create
     */
    includedInBulkChangesets?: boolean;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    name?: string;
    /**
     * When creating, will default to the environment's default namespace, if provided
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    namespace?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    pagerdutyIntegration?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {number}
     * @memberof SherlockChartReleaseV3Create
     */
    port?: number;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    protocol?: string;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3Create
     */
    subdomain?: string;
}


/**
 * @export
 */
export const SherlockChartReleaseV3CreateAppVersionResolverEnum = {
    Branch: 'branch',
    Commit: 'commit',
    Exact: 'exact',
    Follow: 'follow',
    None: 'none'
} as const;
export type SherlockChartReleaseV3CreateAppVersionResolverEnum = typeof SherlockChartReleaseV3CreateAppVersionResolverEnum[keyof typeof SherlockChartReleaseV3CreateAppVersionResolverEnum];

/**
 * @export
 */
export const SherlockChartReleaseV3CreateChartVersionResolverEnum = {
    Latest: 'latest',
    Exact: 'exact',
    Follow: 'follow'
} as const;
export type SherlockChartReleaseV3CreateChartVersionResolverEnum = typeof SherlockChartReleaseV3CreateChartVersionResolverEnum[keyof typeof SherlockChartReleaseV3CreateChartVersionResolverEnum];


/**
 * Check if a given object implements the SherlockChartReleaseV3Create interface.
 */
export function instanceOfSherlockChartReleaseV3Create(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SherlockChartReleaseV3CreateFromJSON(json: any): SherlockChartReleaseV3Create {
    return SherlockChartReleaseV3CreateFromJSONTyped(json, false);
}

export function SherlockChartReleaseV3CreateFromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartReleaseV3Create {
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
        'helmfileRefEnabled': !exists(json, 'helmfileRefEnabled') ? undefined : json['helmfileRefEnabled'],
        'includedInBulkChangesets': !exists(json, 'includedInBulkChangesets') ? undefined : json['includedInBulkChangesets'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'namespace': !exists(json, 'namespace') ? undefined : json['namespace'],
        'pagerdutyIntegration': !exists(json, 'pagerdutyIntegration') ? undefined : json['pagerdutyIntegration'],
        'port': !exists(json, 'port') ? undefined : json['port'],
        'protocol': !exists(json, 'protocol') ? undefined : json['protocol'],
        'subdomain': !exists(json, 'subdomain') ? undefined : json['subdomain'],
    };
}

export function SherlockChartReleaseV3CreateToJSON(value?: SherlockChartReleaseV3Create | null): any {
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
        'helmfileRefEnabled': value.helmfileRefEnabled,
        'includedInBulkChangesets': value.includedInBulkChangesets,
        'name': value.name,
        'namespace': value.namespace,
        'pagerdutyIntegration': value.pagerdutyIntegration,
        'port': value.port,
        'protocol': value.protocol,
        'subdomain': value.subdomain,
    };
}
