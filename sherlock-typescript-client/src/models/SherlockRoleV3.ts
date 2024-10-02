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
import type { SherlockRoleAssignmentV3 } from './SherlockRoleAssignmentV3';
import {
    SherlockRoleAssignmentV3FromJSON,
    SherlockRoleAssignmentV3FromJSONTyped,
    SherlockRoleAssignmentV3ToJSON,
    SherlockRoleAssignmentV3ToJSONTyped,
} from './SherlockRoleAssignmentV3';

/**
 * 
 * @export
 * @interface SherlockRoleV3
 */
export interface SherlockRoleV3 {
    /**
     * 
     * @type {Array<SherlockRoleAssignmentV3>}
     * @memberof SherlockRoleV3
     */
    assignments?: Array<SherlockRoleAssignmentV3>;
    /**
     * When true, Sherlock will automatically assign all users to this role who do not already have a role assignment
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    autoAssignAllUsers?: boolean;
    /**
     * 
     * @type {number}
     * @memberof SherlockRoleV3
     */
    canBeGlassBrokenByRole?: number;
    /**
     * 
     * @type {object}
     * @memberof SherlockRoleV3
     */
    canBeGlassBrokenByRoleInfo?: object;
    /**
     * 
     * @type {Date}
     * @memberof SherlockRoleV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    defaultGlassBreakDuration?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsBroadInstituteGroup?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    grantsDevAzureAccount?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    grantsDevAzureDirectoryRoles?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsDevAzureGroup?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsDevFirecloudFolderOwner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsDevFirecloudGroup?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    grantsProdAzureAccount?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    grantsProdAzureDirectoryRoles?: boolean;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsProdAzureGroup?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsProdFirecloudFolderOwner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsProdFirecloudGroup?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsQaFirecloudFolderOwner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    grantsQaFirecloudGroup?: string;
    /**
     * 
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    grantsSherlockSuperAdmin?: boolean;
    /**
     * 
     * @type {number}
     * @memberof SherlockRoleV3
     */
    id?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockRoleV3
     */
    name?: string;
    /**
     * When true, the "suspended" field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field
     * @type {boolean}
     * @memberof SherlockRoleV3
     */
    suspendNonSuitableUsers?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof SherlockRoleV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockRoleV3 interface.
 */
export function instanceOfSherlockRoleV3(value: object): value is SherlockRoleV3 {
    return true;
}

export function SherlockRoleV3FromJSON(json: any): SherlockRoleV3 {
    return SherlockRoleV3FromJSONTyped(json, false);
}

export function SherlockRoleV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockRoleV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'assignments': json['assignments'] == null ? undefined : ((json['assignments'] as Array<any>).map(SherlockRoleAssignmentV3FromJSON)),
        'autoAssignAllUsers': json['autoAssignAllUsers'] == null ? undefined : json['autoAssignAllUsers'],
        'canBeGlassBrokenByRole': json['canBeGlassBrokenByRole'] == null ? undefined : json['canBeGlassBrokenByRole'],
        'canBeGlassBrokenByRoleInfo': json['canBeGlassBrokenByRoleInfo'] == null ? undefined : json['canBeGlassBrokenByRoleInfo'],
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'defaultGlassBreakDuration': json['defaultGlassBreakDuration'] == null ? undefined : json['defaultGlassBreakDuration'],
        'grantsBroadInstituteGroup': json['grantsBroadInstituteGroup'] == null ? undefined : json['grantsBroadInstituteGroup'],
        'grantsDevAzureAccount': json['grantsDevAzureAccount'] == null ? undefined : json['grantsDevAzureAccount'],
        'grantsDevAzureDirectoryRoles': json['grantsDevAzureDirectoryRoles'] == null ? undefined : json['grantsDevAzureDirectoryRoles'],
        'grantsDevAzureGroup': json['grantsDevAzureGroup'] == null ? undefined : json['grantsDevAzureGroup'],
        'grantsDevFirecloudFolderOwner': json['grantsDevFirecloudFolderOwner'] == null ? undefined : json['grantsDevFirecloudFolderOwner'],
        'grantsDevFirecloudGroup': json['grantsDevFirecloudGroup'] == null ? undefined : json['grantsDevFirecloudGroup'],
        'grantsProdAzureAccount': json['grantsProdAzureAccount'] == null ? undefined : json['grantsProdAzureAccount'],
        'grantsProdAzureDirectoryRoles': json['grantsProdAzureDirectoryRoles'] == null ? undefined : json['grantsProdAzureDirectoryRoles'],
        'grantsProdAzureGroup': json['grantsProdAzureGroup'] == null ? undefined : json['grantsProdAzureGroup'],
        'grantsProdFirecloudFolderOwner': json['grantsProdFirecloudFolderOwner'] == null ? undefined : json['grantsProdFirecloudFolderOwner'],
        'grantsProdFirecloudGroup': json['grantsProdFirecloudGroup'] == null ? undefined : json['grantsProdFirecloudGroup'],
        'grantsQaFirecloudFolderOwner': json['grantsQaFirecloudFolderOwner'] == null ? undefined : json['grantsQaFirecloudFolderOwner'],
        'grantsQaFirecloudGroup': json['grantsQaFirecloudGroup'] == null ? undefined : json['grantsQaFirecloudGroup'],
        'grantsSherlockSuperAdmin': json['grantsSherlockSuperAdmin'] == null ? undefined : json['grantsSherlockSuperAdmin'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'suspendNonSuitableUsers': json['suspendNonSuitableUsers'] == null ? undefined : json['suspendNonSuitableUsers'],
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

  export function SherlockRoleV3ToJSON(json: any): SherlockRoleV3 {
      return SherlockRoleV3ToJSONTyped(json, false);
  }

  export function SherlockRoleV3ToJSONTyped(value?: SherlockRoleV3 | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'assignments': value['assignments'] == null ? undefined : ((value['assignments'] as Array<any>).map(SherlockRoleAssignmentV3ToJSON)),
        'autoAssignAllUsers': value['autoAssignAllUsers'],
        'canBeGlassBrokenByRole': value['canBeGlassBrokenByRole'],
        'canBeGlassBrokenByRoleInfo': value['canBeGlassBrokenByRoleInfo'],
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'defaultGlassBreakDuration': value['defaultGlassBreakDuration'],
        'grantsBroadInstituteGroup': value['grantsBroadInstituteGroup'],
        'grantsDevAzureAccount': value['grantsDevAzureAccount'],
        'grantsDevAzureDirectoryRoles': value['grantsDevAzureDirectoryRoles'],
        'grantsDevAzureGroup': value['grantsDevAzureGroup'],
        'grantsDevFirecloudFolderOwner': value['grantsDevFirecloudFolderOwner'],
        'grantsDevFirecloudGroup': value['grantsDevFirecloudGroup'],
        'grantsProdAzureAccount': value['grantsProdAzureAccount'],
        'grantsProdAzureDirectoryRoles': value['grantsProdAzureDirectoryRoles'],
        'grantsProdAzureGroup': value['grantsProdAzureGroup'],
        'grantsProdFirecloudFolderOwner': value['grantsProdFirecloudFolderOwner'],
        'grantsProdFirecloudGroup': value['grantsProdFirecloudGroup'],
        'grantsQaFirecloudFolderOwner': value['grantsQaFirecloudFolderOwner'],
        'grantsQaFirecloudGroup': value['grantsQaFirecloudGroup'],
        'grantsSherlockSuperAdmin': value['grantsSherlockSuperAdmin'],
        'id': value['id'],
        'name': value['name'],
        'suspendNonSuitableUsers': value['suspendNonSuitableUsers'],
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

