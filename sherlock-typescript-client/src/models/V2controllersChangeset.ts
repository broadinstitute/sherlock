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
import type { V2controllersChartRelease } from './V2controllersChartRelease';
import {
    V2controllersChartReleaseFromJSON,
    V2controllersChartReleaseFromJSONTyped,
    V2controllersChartReleaseToJSON,
} from './V2controllersChartRelease';
import type { V2controllersChartVersion } from './V2controllersChartVersion';
import {
    V2controllersChartVersionFromJSON,
    V2controllersChartVersionFromJSONTyped,
    V2controllersChartVersionToJSON,
} from './V2controllersChartVersion';
import type { V2controllersCiIdentifier } from './V2controllersCiIdentifier';
import {
    V2controllersCiIdentifierFromJSON,
    V2controllersCiIdentifierFromJSONTyped,
    V2controllersCiIdentifierToJSON,
} from './V2controllersCiIdentifier';
import type { V2controllersUser } from './V2controllersUser';
import {
    V2controllersUserFromJSON,
    V2controllersUserFromJSONTyped,
    V2controllersUserToJSON,
} from './V2controllersUser';

/**
 * 
 * @export
 * @interface V2controllersChangeset
 */
export interface V2controllersChangeset {
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    appliedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    appliedBy?: string;
    /**
     * 
     * @type {V2controllersUser}
     * @memberof V2controllersChangeset
     */
    appliedByInfo?: V2controllersUser;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    chartRelease?: string;
    /**
     * 
     * @type {V2controllersChartRelease}
     * @memberof V2controllersChangeset
     */
    chartReleaseInfo?: V2controllersChartRelease;
    /**
     * 
     * @type {V2controllersCiIdentifier}
     * @memberof V2controllersChangeset
     */
    ciIdentifier?: V2controllersCiIdentifier;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    createdAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionReference?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromAppVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromChartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromChartVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromChartVersionReference?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromChartVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromFirecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    fromHelmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersChangeset
     */
    fromHelmfileRefEnabled?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    fromResolvedAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof V2controllersChangeset
     */
    id?: number;
    /**
     * 
     * @type {Array<V2controllersAppVersion>}
     * @memberof V2controllersChangeset
     */
    newAppVersions?: Array<V2controllersAppVersion>;
    /**
     * 
     * @type {Array<V2controllersChartVersion>}
     * @memberof V2controllersChangeset
     */
    newChartVersions?: Array<V2controllersChartVersion>;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    plannedBy?: string;
    /**
     * 
     * @type {V2controllersUser}
     * @memberof V2controllersChangeset
     */
    plannedByInfo?: V2controllersUser;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    supersededAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionBranch?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionCommit?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionReference?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toAppVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toChartVersionExact?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toChartVersionFollowChartRelease?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toChartVersionReference?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toChartVersionResolver?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toFirecloudDevelopRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V2controllersChangeset
     */
    toHelmfileRef?: string;
    /**
     * 
     * @type {boolean}
     * @memberof V2controllersChangeset
     */
    toHelmfileRefEnabled?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    toResolvedAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V2controllersChangeset
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the V2controllersChangeset interface.
 */
export function instanceOfV2controllersChangeset(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V2controllersChangesetFromJSON(json: any): V2controllersChangeset {
    return V2controllersChangesetFromJSONTyped(json, false);
}

export function V2controllersChangesetFromJSONTyped(json: any, ignoreDiscriminator: boolean): V2controllersChangeset {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appliedAt': !exists(json, 'appliedAt') ? undefined : (new Date(json['appliedAt'])),
        'appliedBy': !exists(json, 'appliedBy') ? undefined : json['appliedBy'],
        'appliedByInfo': !exists(json, 'appliedByInfo') ? undefined : V2controllersUserFromJSON(json['appliedByInfo']),
        'chartRelease': !exists(json, 'chartRelease') ? undefined : json['chartRelease'],
        'chartReleaseInfo': !exists(json, 'chartReleaseInfo') ? undefined : V2controllersChartReleaseFromJSON(json['chartReleaseInfo']),
        'ciIdentifier': !exists(json, 'ciIdentifier') ? undefined : V2controllersCiIdentifierFromJSON(json['ciIdentifier']),
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'fromAppVersionBranch': !exists(json, 'fromAppVersionBranch') ? undefined : json['fromAppVersionBranch'],
        'fromAppVersionCommit': !exists(json, 'fromAppVersionCommit') ? undefined : json['fromAppVersionCommit'],
        'fromAppVersionExact': !exists(json, 'fromAppVersionExact') ? undefined : json['fromAppVersionExact'],
        'fromAppVersionFollowChartRelease': !exists(json, 'fromAppVersionFollowChartRelease') ? undefined : json['fromAppVersionFollowChartRelease'],
        'fromAppVersionReference': !exists(json, 'fromAppVersionReference') ? undefined : json['fromAppVersionReference'],
        'fromAppVersionResolver': !exists(json, 'fromAppVersionResolver') ? undefined : json['fromAppVersionResolver'],
        'fromChartVersionExact': !exists(json, 'fromChartVersionExact') ? undefined : json['fromChartVersionExact'],
        'fromChartVersionFollowChartRelease': !exists(json, 'fromChartVersionFollowChartRelease') ? undefined : json['fromChartVersionFollowChartRelease'],
        'fromChartVersionReference': !exists(json, 'fromChartVersionReference') ? undefined : json['fromChartVersionReference'],
        'fromChartVersionResolver': !exists(json, 'fromChartVersionResolver') ? undefined : json['fromChartVersionResolver'],
        'fromFirecloudDevelopRef': !exists(json, 'fromFirecloudDevelopRef') ? undefined : json['fromFirecloudDevelopRef'],
        'fromHelmfileRef': !exists(json, 'fromHelmfileRef') ? undefined : json['fromHelmfileRef'],
        'fromHelmfileRefEnabled': !exists(json, 'fromHelmfileRefEnabled') ? undefined : json['fromHelmfileRefEnabled'],
        'fromResolvedAt': !exists(json, 'fromResolvedAt') ? undefined : (new Date(json['fromResolvedAt'])),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'newAppVersions': !exists(json, 'newAppVersions') ? undefined : ((json['newAppVersions'] as Array<any>).map(V2controllersAppVersionFromJSON)),
        'newChartVersions': !exists(json, 'newChartVersions') ? undefined : ((json['newChartVersions'] as Array<any>).map(V2controllersChartVersionFromJSON)),
        'plannedBy': !exists(json, 'plannedBy') ? undefined : json['plannedBy'],
        'plannedByInfo': !exists(json, 'plannedByInfo') ? undefined : V2controllersUserFromJSON(json['plannedByInfo']),
        'supersededAt': !exists(json, 'supersededAt') ? undefined : (new Date(json['supersededAt'])),
        'toAppVersionBranch': !exists(json, 'toAppVersionBranch') ? undefined : json['toAppVersionBranch'],
        'toAppVersionCommit': !exists(json, 'toAppVersionCommit') ? undefined : json['toAppVersionCommit'],
        'toAppVersionExact': !exists(json, 'toAppVersionExact') ? undefined : json['toAppVersionExact'],
        'toAppVersionFollowChartRelease': !exists(json, 'toAppVersionFollowChartRelease') ? undefined : json['toAppVersionFollowChartRelease'],
        'toAppVersionReference': !exists(json, 'toAppVersionReference') ? undefined : json['toAppVersionReference'],
        'toAppVersionResolver': !exists(json, 'toAppVersionResolver') ? undefined : json['toAppVersionResolver'],
        'toChartVersionExact': !exists(json, 'toChartVersionExact') ? undefined : json['toChartVersionExact'],
        'toChartVersionFollowChartRelease': !exists(json, 'toChartVersionFollowChartRelease') ? undefined : json['toChartVersionFollowChartRelease'],
        'toChartVersionReference': !exists(json, 'toChartVersionReference') ? undefined : json['toChartVersionReference'],
        'toChartVersionResolver': !exists(json, 'toChartVersionResolver') ? undefined : json['toChartVersionResolver'],
        'toFirecloudDevelopRef': !exists(json, 'toFirecloudDevelopRef') ? undefined : json['toFirecloudDevelopRef'],
        'toHelmfileRef': !exists(json, 'toHelmfileRef') ? undefined : json['toHelmfileRef'],
        'toHelmfileRefEnabled': !exists(json, 'toHelmfileRefEnabled') ? undefined : json['toHelmfileRefEnabled'],
        'toResolvedAt': !exists(json, 'toResolvedAt') ? undefined : (new Date(json['toResolvedAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
    };
}

export function V2controllersChangesetToJSON(value?: V2controllersChangeset | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appliedAt': value.appliedAt === undefined ? undefined : (value.appliedAt.toISOString()),
        'appliedBy': value.appliedBy,
        'appliedByInfo': V2controllersUserToJSON(value.appliedByInfo),
        'chartRelease': value.chartRelease,
        'chartReleaseInfo': V2controllersChartReleaseToJSON(value.chartReleaseInfo),
        'ciIdentifier': V2controllersCiIdentifierToJSON(value.ciIdentifier),
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'fromAppVersionBranch': value.fromAppVersionBranch,
        'fromAppVersionCommit': value.fromAppVersionCommit,
        'fromAppVersionExact': value.fromAppVersionExact,
        'fromAppVersionFollowChartRelease': value.fromAppVersionFollowChartRelease,
        'fromAppVersionReference': value.fromAppVersionReference,
        'fromAppVersionResolver': value.fromAppVersionResolver,
        'fromChartVersionExact': value.fromChartVersionExact,
        'fromChartVersionFollowChartRelease': value.fromChartVersionFollowChartRelease,
        'fromChartVersionReference': value.fromChartVersionReference,
        'fromChartVersionResolver': value.fromChartVersionResolver,
        'fromFirecloudDevelopRef': value.fromFirecloudDevelopRef,
        'fromHelmfileRef': value.fromHelmfileRef,
        'fromHelmfileRefEnabled': value.fromHelmfileRefEnabled,
        'fromResolvedAt': value.fromResolvedAt === undefined ? undefined : (value.fromResolvedAt.toISOString()),
        'id': value.id,
        'newAppVersions': value.newAppVersions === undefined ? undefined : ((value.newAppVersions as Array<any>).map(V2controllersAppVersionToJSON)),
        'newChartVersions': value.newChartVersions === undefined ? undefined : ((value.newChartVersions as Array<any>).map(V2controllersChartVersionToJSON)),
        'plannedBy': value.plannedBy,
        'plannedByInfo': V2controllersUserToJSON(value.plannedByInfo),
        'supersededAt': value.supersededAt === undefined ? undefined : (value.supersededAt.toISOString()),
        'toAppVersionBranch': value.toAppVersionBranch,
        'toAppVersionCommit': value.toAppVersionCommit,
        'toAppVersionExact': value.toAppVersionExact,
        'toAppVersionFollowChartRelease': value.toAppVersionFollowChartRelease,
        'toAppVersionReference': value.toAppVersionReference,
        'toAppVersionResolver': value.toAppVersionResolver,
        'toChartVersionExact': value.toChartVersionExact,
        'toChartVersionFollowChartRelease': value.toChartVersionFollowChartRelease,
        'toChartVersionReference': value.toChartVersionReference,
        'toChartVersionResolver': value.toChartVersionResolver,
        'toFirecloudDevelopRef': value.toFirecloudDevelopRef,
        'toHelmfileRef': value.toHelmfileRef,
        'toHelmfileRefEnabled': value.toHelmfileRefEnabled,
        'toResolvedAt': value.toResolvedAt === undefined ? undefined : (value.toResolvedAt.toISOString()),
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
    };
}

