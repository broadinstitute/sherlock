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

import { mapValues } from '../runtime';
import type { SherlockClusterV3 } from './SherlockClusterV3';
import {
    SherlockClusterV3FromJSON,
    SherlockClusterV3FromJSONTyped,
    SherlockClusterV3ToJSON,
} from './SherlockClusterV3';
import type { SherlockChartVersionV3 } from './SherlockChartVersionV3';
import {
    SherlockChartVersionV3FromJSON,
    SherlockChartVersionV3FromJSONTyped,
    SherlockChartVersionV3ToJSON,
} from './SherlockChartVersionV3';
import type { SherlockChartV3 } from './SherlockChartV3';
import {
    SherlockChartV3FromJSON,
    SherlockChartV3FromJSONTyped,
    SherlockChartV3ToJSON,
} from './SherlockChartV3';
import type { SherlockAppVersionV3 } from './SherlockAppVersionV3';
import {
    SherlockAppVersionV3FromJSON,
    SherlockAppVersionV3FromJSONTyped,
    SherlockAppVersionV3ToJSON,
} from './SherlockAppVersionV3';
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
} from './SherlockCiIdentifierV3';
import type { SherlockEnvironmentV3 } from './SherlockEnvironmentV3';
import {
    SherlockEnvironmentV3FromJSON,
    SherlockEnvironmentV3FromJSONTyped,
    SherlockEnvironmentV3ToJSON,
} from './SherlockEnvironmentV3';
import type { SherlockPagerdutyIntegrationV3 } from './SherlockPagerdutyIntegrationV3';
import {
    SherlockPagerdutyIntegrationV3FromJSON,
    SherlockPagerdutyIntegrationV3FromJSONTyped,
    SherlockPagerdutyIntegrationV3ToJSON,
} from './SherlockPagerdutyIntegrationV3';

/**
 * 
 * @export
 * @interface SherlockChartReleaseV3
 */
export interface SherlockChartReleaseV3 {
    /**
     * When creating, will default to the app's mainline branch if no other app version info is present
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionFollowChartRelease?: string;
    /**
     * 
     * @type {SherlockAppVersionV3}
     * @memberof SherlockChartReleaseV3
     */
    appVersionInfo?: SherlockAppVersionV3;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionReference?: string;
    /**
     * // When creating, will default to automatically reference any provided app version fields
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    appVersionResolver?: SherlockChartReleaseV3AppVersionResolverEnum;
    /**
     * Required when creating
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    chart?: string;
    /**
     * 
     * @type {SherlockChartV3}
     * @memberof SherlockChartReleaseV3
     */
    chartInfo?: SherlockChartV3;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    chartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    chartVersionFollowChartRelease?: string;
    /**
     * 
     * @type {SherlockChartVersionV3}
     * @memberof SherlockChartReleaseV3
     */
    chartVersionInfo?: SherlockChartVersionV3;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    chartVersionReference?: string;
    /**
     * When creating, will default to automatically reference any provided chart version
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    chartVersionResolver?: SherlockChartReleaseV3ChartVersionResolverEnum;
    /**
     * 
     * @type {SherlockCiIdentifierV3}
     * @memberof SherlockChartReleaseV3
     */
    ciIdentifier?: SherlockCiIdentifierV3;
    /**
     * When creating, will default the environment's default cluster, if provided. Either this or environment must be provided.
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    cluster?: string;
    /**
     * 
     * @type {SherlockClusterV3}
     * @memberof SherlockChartReleaseV3
     */
    clusterInfo?: SherlockClusterV3;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartReleaseV3
     */
    createdAt?: Date;
    /**
     * Calculated field
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    destinationType?: string;
    /**
     * Either this or cluster must be provided.
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    environment?: string;
    /**
     * 
     * @type {SherlockEnvironmentV3}
     * @memberof SherlockChartReleaseV3
     */
    environmentInfo?: SherlockEnvironmentV3;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    helmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartReleaseV3
     */
    helmfileRefEnabled?: boolean;
    /**
     * 
     * @type {number}
     * @memberof SherlockChartReleaseV3
     */
    id?: number;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockChartReleaseV3
     */
    includedInBulkChangesets?: boolean;
    /**
     * When creating, will be calculated if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    name?: string;
    /**
     * When creating, will default to the environment's default namespace, if provided
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    namespace?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    pagerdutyIntegration?: string;
    /**
     * 
     * @type {SherlockPagerdutyIntegrationV3}
     * @memberof SherlockChartReleaseV3
     */
    pagerdutyIntegrationInfo?: SherlockPagerdutyIntegrationV3;
    /**
     * When creating, will use the chart's default if left empty
     * @type {number}
     * @memberof SherlockChartReleaseV3
     */
    port?: number;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    protocol?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartReleaseV3
     */
    resolvedAt?: Date;
    /**
     * When creating, will use the chart's default if left empty
     * @type {string}
     * @memberof SherlockChartReleaseV3
     */
    subdomain?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockChartReleaseV3
     */
    updatedAt?: Date;
}


/**
 * @export
 */
export const SherlockChartReleaseV3AppVersionResolverEnum = {
    Branch: 'branch',
    Commit: 'commit',
    Exact: 'exact',
    Follow: 'follow',
    None: 'none'
} as const;
export type SherlockChartReleaseV3AppVersionResolverEnum = typeof SherlockChartReleaseV3AppVersionResolverEnum[keyof typeof SherlockChartReleaseV3AppVersionResolverEnum];

/**
 * @export
 */
export const SherlockChartReleaseV3ChartVersionResolverEnum = {
    Latest: 'latest',
    Exact: 'exact',
    Follow: 'follow'
} as const;
export type SherlockChartReleaseV3ChartVersionResolverEnum = typeof SherlockChartReleaseV3ChartVersionResolverEnum[keyof typeof SherlockChartReleaseV3ChartVersionResolverEnum];


/**
 * Check if a given object implements the SherlockChartReleaseV3 interface.
 */
export function instanceOfSherlockChartReleaseV3(value: object): value is SherlockChartReleaseV3 {
    return true;
}

export function SherlockChartReleaseV3FromJSON(json: any): SherlockChartReleaseV3 {
    return SherlockChartReleaseV3FromJSONTyped(json, false);
}

export function SherlockChartReleaseV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockChartReleaseV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'appVersionBranch': json['appVersionBranch'] == null ? undefined : json['appVersionBranch'],
        'appVersionCommit': json['appVersionCommit'] == null ? undefined : json['appVersionCommit'],
        'appVersionExact': json['appVersionExact'] == null ? undefined : json['appVersionExact'],
        'appVersionFollowChartRelease': json['appVersionFollowChartRelease'] == null ? undefined : json['appVersionFollowChartRelease'],
        'appVersionInfo': json['appVersionInfo'] == null ? undefined : SherlockAppVersionV3FromJSON(json['appVersionInfo']),
        'appVersionReference': json['appVersionReference'] == null ? undefined : json['appVersionReference'],
        'appVersionResolver': json['appVersionResolver'] == null ? undefined : json['appVersionResolver'],
        'chart': json['chart'] == null ? undefined : json['chart'],
        'chartInfo': json['chartInfo'] == null ? undefined : SherlockChartV3FromJSON(json['chartInfo']),
        'chartVersionExact': json['chartVersionExact'] == null ? undefined : json['chartVersionExact'],
        'chartVersionFollowChartRelease': json['chartVersionFollowChartRelease'] == null ? undefined : json['chartVersionFollowChartRelease'],
        'chartVersionInfo': json['chartVersionInfo'] == null ? undefined : SherlockChartVersionV3FromJSON(json['chartVersionInfo']),
        'chartVersionReference': json['chartVersionReference'] == null ? undefined : json['chartVersionReference'],
        'chartVersionResolver': json['chartVersionResolver'] == null ? undefined : json['chartVersionResolver'],
        'ciIdentifier': json['ciIdentifier'] == null ? undefined : SherlockCiIdentifierV3FromJSON(json['ciIdentifier']),
        'cluster': json['cluster'] == null ? undefined : json['cluster'],
        'clusterInfo': json['clusterInfo'] == null ? undefined : SherlockClusterV3FromJSON(json['clusterInfo']),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'destinationType': json['destinationType'] == null ? undefined : json['destinationType'],
        'environment': json['environment'] == null ? undefined : json['environment'],
        'environmentInfo': json['environmentInfo'] == null ? undefined : SherlockEnvironmentV3FromJSON(json['environmentInfo']),
        'helmfileRef': json['helmfileRef'] == null ? undefined : json['helmfileRef'],
        'helmfileRefEnabled': json['helmfileRefEnabled'] == null ? undefined : json['helmfileRefEnabled'],
        'id': json['id'] == null ? undefined : json['id'],
        'includedInBulkChangesets': json['includedInBulkChangesets'] == null ? undefined : json['includedInBulkChangesets'],
        'name': json['name'] == null ? undefined : json['name'],
        'namespace': json['namespace'] == null ? undefined : json['namespace'],
        'pagerdutyIntegration': json['pagerdutyIntegration'] == null ? undefined : json['pagerdutyIntegration'],
        'pagerdutyIntegrationInfo': json['pagerdutyIntegrationInfo'] == null ? undefined : SherlockPagerdutyIntegrationV3FromJSON(json['pagerdutyIntegrationInfo']),
        'port': json['port'] == null ? undefined : json['port'],
        'protocol': json['protocol'] == null ? undefined : json['protocol'],
        'resolvedAt': json['resolvedAt'] == null ? undefined : (new Date(json['resolvedAt'])),
        'subdomain': json['subdomain'] == null ? undefined : json['subdomain'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

export function SherlockChartReleaseV3ToJSON(value?: SherlockChartReleaseV3 | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'appVersionBranch': value['appVersionBranch'],
        'appVersionCommit': value['appVersionCommit'],
        'appVersionExact': value['appVersionExact'],
        'appVersionFollowChartRelease': value['appVersionFollowChartRelease'],
        'appVersionInfo': SherlockAppVersionV3ToJSON(value['appVersionInfo']),
        'appVersionReference': value['appVersionReference'],
        'appVersionResolver': value['appVersionResolver'],
        'chart': value['chart'],
        'chartInfo': SherlockChartV3ToJSON(value['chartInfo']),
        'chartVersionExact': value['chartVersionExact'],
        'chartVersionFollowChartRelease': value['chartVersionFollowChartRelease'],
        'chartVersionInfo': SherlockChartVersionV3ToJSON(value['chartVersionInfo']),
        'chartVersionReference': value['chartVersionReference'],
        'chartVersionResolver': value['chartVersionResolver'],
        'ciIdentifier': SherlockCiIdentifierV3ToJSON(value['ciIdentifier']),
        'cluster': value['cluster'],
        'clusterInfo': SherlockClusterV3ToJSON(value['clusterInfo']),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'destinationType': value['destinationType'],
        'environment': value['environment'],
        'environmentInfo': SherlockEnvironmentV3ToJSON(value['environmentInfo']),
        'helmfileRef': value['helmfileRef'],
        'helmfileRefEnabled': value['helmfileRefEnabled'],
        'id': value['id'],
        'includedInBulkChangesets': value['includedInBulkChangesets'],
        'name': value['name'],
        'namespace': value['namespace'],
        'pagerdutyIntegration': value['pagerdutyIntegration'],
        'pagerdutyIntegrationInfo': SherlockPagerdutyIntegrationV3ToJSON(value['pagerdutyIntegrationInfo']),
        'port': value['port'],
        'protocol': value['protocol'],
        'resolvedAt': value['resolvedAt'] == null ? undefined : ((value['resolvedAt']).toISOString()),
        'subdomain': value['subdomain'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

