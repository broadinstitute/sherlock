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
import type { SherlockCiIdentifierV3 } from './SherlockCiIdentifierV3';
import {
    SherlockCiIdentifierV3FromJSON,
    SherlockCiIdentifierV3FromJSONTyped,
    SherlockCiIdentifierV3ToJSON,
    SherlockCiIdentifierV3ToJSONTyped,
} from './SherlockCiIdentifierV3';

/**
 * 
 * @export
 * @interface SherlockCiRunV3
 */
export interface SherlockCiRunV3 {
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    argoWorkflowsName?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    argoWorkflowsNamespace?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    argoWorkflowsTemplate?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockCiRunV3
     */
    createdAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof SherlockCiRunV3
     */
    githubActionsAttemptNumber?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    githubActionsOwner?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    githubActionsRepo?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockCiRunV3
     */
    githubActionsRunID?: number;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    githubActionsWorkflowPath?: string;
    /**
     * 
     * @type {number}
     * @memberof SherlockCiRunV3
     */
    id?: number;
    /**
     * Slack channels to notify if this CiRun fails. This field is always appended to when mutated.
     * @type {Array<string>}
     * @memberof SherlockCiRunV3
     */
    notifySlackChannelsUponFailure?: Array<string>;
    /**
     * Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields.
     * @type {Array<string>}
     * @memberof SherlockCiRunV3
     */
    notifySlackChannelsUponRetry?: Array<string>;
    /**
     * Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated.
     * @type {Array<string>}
     * @memberof SherlockCiRunV3
     */
    notifySlackChannelsUponSuccess?: Array<string>;
    /**
     * Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:).
     * An empty string is ignored to facilitate calling from GitHub Actions (where it's easier to pass an empty string than not send the field at all).
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    notifySlackCustomIcon?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    platform?: string;
    /**
     * 
     * @type {Array<SherlockCiIdentifierV3>}
     * @memberof SherlockCiRunV3
     */
    relatedResources?: Array<SherlockCiIdentifierV3>;
    /**
     * Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    resourceStatus?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    startedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    status?: string;
    /**
     * 
     * @type {string}
     * @memberof SherlockCiRunV3
     */
    terminalAt?: string;
    /**
     * 
     * @type {Date}
     * @memberof SherlockCiRunV3
     */
    terminationHooksDispatchedAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof SherlockCiRunV3
     */
    updatedAt?: Date;
}

/**
 * Check if a given object implements the SherlockCiRunV3 interface.
 */
export function instanceOfSherlockCiRunV3(value: object): value is SherlockCiRunV3 {
    return true;
}

export function SherlockCiRunV3FromJSON(json: any): SherlockCiRunV3 {
    return SherlockCiRunV3FromJSONTyped(json, false);
}

export function SherlockCiRunV3FromJSONTyped(json: any, ignoreDiscriminator: boolean): SherlockCiRunV3 {
    if (json == null) {
        return json;
    }
    return {
        
        'argoWorkflowsName': json['argoWorkflowsName'] == null ? undefined : json['argoWorkflowsName'],
        'argoWorkflowsNamespace': json['argoWorkflowsNamespace'] == null ? undefined : json['argoWorkflowsNamespace'],
        'argoWorkflowsTemplate': json['argoWorkflowsTemplate'] == null ? undefined : json['argoWorkflowsTemplate'],
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'githubActionsAttemptNumber': json['githubActionsAttemptNumber'] == null ? undefined : json['githubActionsAttemptNumber'],
        'githubActionsOwner': json['githubActionsOwner'] == null ? undefined : json['githubActionsOwner'],
        'githubActionsRepo': json['githubActionsRepo'] == null ? undefined : json['githubActionsRepo'],
        'githubActionsRunID': json['githubActionsRunID'] == null ? undefined : json['githubActionsRunID'],
        'githubActionsWorkflowPath': json['githubActionsWorkflowPath'] == null ? undefined : json['githubActionsWorkflowPath'],
        'id': json['id'] == null ? undefined : json['id'],
        'notifySlackChannelsUponFailure': json['notifySlackChannelsUponFailure'] == null ? undefined : json['notifySlackChannelsUponFailure'],
        'notifySlackChannelsUponRetry': json['notifySlackChannelsUponRetry'] == null ? undefined : json['notifySlackChannelsUponRetry'],
        'notifySlackChannelsUponSuccess': json['notifySlackChannelsUponSuccess'] == null ? undefined : json['notifySlackChannelsUponSuccess'],
        'notifySlackCustomIcon': json['notifySlackCustomIcon'] == null ? undefined : json['notifySlackCustomIcon'],
        'platform': json['platform'] == null ? undefined : json['platform'],
        'relatedResources': json['relatedResources'] == null ? undefined : ((json['relatedResources'] as Array<any>).map(SherlockCiIdentifierV3FromJSON)),
        'resourceStatus': json['resourceStatus'] == null ? undefined : json['resourceStatus'],
        'startedAt': json['startedAt'] == null ? undefined : json['startedAt'],
        'status': json['status'] == null ? undefined : json['status'],
        'terminalAt': json['terminalAt'] == null ? undefined : json['terminalAt'],
        'terminationHooksDispatchedAt': json['terminationHooksDispatchedAt'] == null ? undefined : (new Date(json['terminationHooksDispatchedAt'])),
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
    };
}

  export function SherlockCiRunV3ToJSON(json: any): SherlockCiRunV3 {
      return SherlockCiRunV3ToJSONTyped(json, false);
  }

  export function SherlockCiRunV3ToJSONTyped(value?: SherlockCiRunV3 | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'argoWorkflowsName': value['argoWorkflowsName'],
        'argoWorkflowsNamespace': value['argoWorkflowsNamespace'],
        'argoWorkflowsTemplate': value['argoWorkflowsTemplate'],
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'githubActionsAttemptNumber': value['githubActionsAttemptNumber'],
        'githubActionsOwner': value['githubActionsOwner'],
        'githubActionsRepo': value['githubActionsRepo'],
        'githubActionsRunID': value['githubActionsRunID'],
        'githubActionsWorkflowPath': value['githubActionsWorkflowPath'],
        'id': value['id'],
        'notifySlackChannelsUponFailure': value['notifySlackChannelsUponFailure'],
        'notifySlackChannelsUponRetry': value['notifySlackChannelsUponRetry'],
        'notifySlackChannelsUponSuccess': value['notifySlackChannelsUponSuccess'],
        'notifySlackCustomIcon': value['notifySlackCustomIcon'],
        'platform': value['platform'],
        'relatedResources': value['relatedResources'] == null ? undefined : ((value['relatedResources'] as Array<any>).map(SherlockCiIdentifierV3ToJSON)),
        'resourceStatus': value['resourceStatus'],
        'startedAt': value['startedAt'],
        'status': value['status'],
        'terminalAt': value['terminalAt'],
        'terminationHooksDispatchedAt': value['terminationHooksDispatchedAt'] == null ? undefined : ((value['terminationHooksDispatchedAt']).toISOString()),
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
    };
}

