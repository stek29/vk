import json
import os

#TODO: Hanlde type aliases for types with custom marshallers (boolint)

methods = {}

with open('vk-api-schema/methods.json') as f:
	j = json.load(f)
	for method in j['methods']:
		namespace = method['name'].split('.')[0]
		if namespace not in methods:
			methods[namespace] = []

		methods[namespace].append(method)


def handle_special_caps(s):
	return s.replace('Id', 'ID')\
		.replace('Url', 'URL')\
		.replace('Guid', 'GUID')\
		.replace('Api', 'API')\
		.replace('Uid', 'UID')\
		.replace('Html', 'HTML')\
		.replace('Ip', 'IP')\
		.replace('Https', 'HTTPS')

def goify_namespace(ns):
	return ns.title()

def goify_method(mtd):
	if '.' in mtd:
		mtd = mtd.split('.')[1]

	mtd = mtd[0].upper() + mtd[1:]

	mtd = handle_special_caps(mtd)

	return mtd

NATIVE_TYPES = {
	'boolean': 'bool',
	'integer': 'int',
	'string': 'string',
	'number': 'float32',
}

NATIVE_TYPES_ZERO_VALS = {
	'bool': ('false', """
		var cnv bool
		cnv, err = strconv.ParseBool(string(r))
		{} = {}(cnv)
	"""),
	'int': ('0', """
		var cnv int
		cnv, err = strconv.Atoi(string(r))
		{} = {}(cnv)
	"""),
	'string': ('""', """
		{} = {}(string(r))
	"""),
	'float32': ('float32(0)', """
		var cnv float32
		cnv, err = strconv.ParseFloat(string(r), 32)
		{} = {}(cnv)
	"""),
}

KNOWN_TYPE_REFS = {
	'wall_comment_attachment': 'Attachment',
	'wall_wallpost_attachment': 'Attachment',
	'messages_history_message_attachment': 'Attachment',

	'wall_wallpost_full': 'Post',
	'wall_wallpost': 'Post',
	'wall_wallpost_attached': 'Post',
	'newsfeed_item_wallpost': 'Post',
	'wall_wallpost_to_id': 'Post',

	'wall_wall_comment': 'Comment',
	# XXX: pid is lost
	'photos_comment_xtr_pid': 'Comment',
	

	# FIXME: xtr are lost
	'photos_photo': 'Photo',
	'photos_photo_full': 'Photo',
	'photos_photo_xtr_tag_info': 'Photo',
	'photos_photo_xtr_real_offset': 'Photo',
	'photos_photo_full_xtr_real_offset': 'Photo',

	# TODO: VideoCat{Block,Element}
	# TODO: Video/Photo Albums
	# TODO: apps_app

	'users_user_min': 'User',
	'users_user_full': 'User',
	'users_user': 'User',
	# lost xtrs
	'users_user_full_xtr_type': 'User',
	'users_user_xtr_type': 'User',
	'users_user_xtr_counters': 'User',
	'friends_user_xtr_phone': 'User',

	'groups_group_full': 'Group',
	'groups_group': 'Group',
	# Lost xtr
	'groups_group_xtr_invited_by': 'Group',

	'base_ok_response': 'BoolInt',
	'base_property_exists': 'BoolInt',
	'base_bool_int': 'BoolInt',

	'base_object': 'BaseObject',
	'base_object_with_name': 'BaseObjectWithName',
	'base_country': 'BaseObject',
	'database_faculty': 'BaseObject',
	'database_faculty': 'BaseObject',
	'database_region': 'BaseObject',
	'database_school': 'BaseObject',
	'database_street': 'BaseObject',
	'database_university': 'BaseObject',
	'database_city': 'DatabaseCity',

	'video_video': 'Video',
	'video_video_full': 'Video',

	'docs_doc': 'Document',

	'board_topic_comment': 'CommentBoard',

	# Tag stuff is lost, but this method is deprecated anyway
	'video_video_tag_info': 'Video',

	'audio_audio': 'Audio',
	'audio_audio_full': 'Audio',

	'photos_photo': 'Photo',
	'photos_photo_full': 'Photo',

	'users_crop_photo': 'CropPhoto',

	'messages_chat': 'Chat',

	'messages_message': 'Message',
	'messages_conversation': 'Conversation',

	'pages_wikipage': 'Page',
	'pages_wikipage_full': 'Page',

	'base_image': 'BaseImage',

	'gifts_gift': 'Gift',

	'fave_faves_link': 'MiniLink',

	'market_market_album': 'MarketAlbum',
	'market_market_item': 'MarketItem',
	'market_market_item_full': 'MarketItem',
	'market_market_category': 'MarketCategory',

	'board_topic': 'BoardTopic',

	'board_topic_poll': 'Poll',
	'polls_poll': 'Poll',

	'places_place_min': 'Place',
	'places_place_full': 'Place',

	'groups_group_category': 'Category',
	'groups_group_public_category_list': 'Category',

	'stats_wallpost_stat': 'WallpostStats',

	'stories_story': 'Story',
	'notes_note': 'Note',
}

FORCE_CODE_GENERATE = {
	'messages_conversation_member',

	'groups_group_category_full',

	'video_save_result',

	'account_lookup_result', 'account_other_contact', 'account_user_xtr_contact',

	'messages_messages_array',
	'users_users_array',
	'groups_groups_array',
	'status_status',
	'messages_history_attachment',

	'messages_conversation_with_message',

	'docs_doc_types',

	'friends_friends_list',

	'pages_wikipage_version',

	'account_info', 'account_user_settings',
	'account_account_counters',

	'account_name_request', 'account_name_request_status',
	'account_offer',
	'account_push_settings', 'account_push_conversations', 'account_push_conversations_item',
	'account_push_params', 'account_push_params', 'account_push_params_mode', 'account_push_params_settings',
	'account_onoff_options',
	
	'base_sex',

	'ads_account', 'ads_account_type', 'ads_access_role',

	'search_hint', 'search_hint_type', 'search_hint_section',

	'secure_sms_notification', 'secure_level', 'secure_transaction', 'secure_token_checked',

	'utils_domain_resolved', 'utils_domain_resolved_type',

	'utils_short_link', 'utils_last_shortened_link', 'utils_link_stats', 'utils_link_stats_extended',

	'utils_stats', 'utils_stats_extended', 'utils_stats_sex_age', 'utils_stats_country', 'utils_stats_city',

	'utils_link_checked', 'utils_link_checked_status',

	'orders_order', 'orders_amount', 'orders_amount_item',

	'groups_callback_settings', 
	'groups_long_poll_settings',
	'groups_long_poll_events',

	'base_upload_server',

	'groups_long_poll_server',

	'groups_owner_xtr_ban_info', 'groups_owner_xtr_ban_info_type',
	'groups_ban_info', 'groups_ban_info_reason',
	'messages_longpoll_params',
	'messages_longpoll_messages',
	'messages_last_activity',
	'messages_email',

	'board_default_order',

	'polls_voters', 'polls_voters_users',

	'groups_group_settings',
	'places_checkin',

	'leads_checked',
	'leads_complete',
	'leads_entry',
	'leads_lead',
	'leads_start',
	'leads_checked_result',
	'leads_lead_days',

	'apps_leaderboard',
	'places_types',
	'stats_wallpost_stat',

	'friends_friend_status', 'friends_friend_status_status',

	'stories_story_stats', 'stories_story_stats_stat', 'stories_story_stats_state',

	'photos_photo_upload',
}

def goify_field_name(name):
	name = handle_special_caps(''.join(x.title() for x in name.split('_')))

	# Thanks for 2fa_required, VK!
	if name[0].isdigit():
		name = 'X' + name

	return name

def goify_field(field):
	# TODO: Enums
	go_name = goify_field_name(field['name'])

	is_required = field.get('required', False)
	url_tag = field['name']
	if not is_required:
		url_tag = url_tag + ",omitempty"
	url_tag = '`url:"{}"`'.format(url_tag)

	t = field['type']
	is_array = t == 'array'
	if is_array:
		native = {
			'integer': 'CSVIntSlice',
			'string': 'CSVStringSlice',
		}
		t = field['items']['type']

		go_type = native[t]
	else:
		go_type = NATIVE_TYPES[t]

	return '\t{} {} {}'.format(go_name, go_type, url_tag)

def goify_object(o):
	if 'properties' in o:
		lines = ['struct {']

		for k, v in o['properties'].items():
			go_name = goify_field_name(k)
			typedef, desc = goify_type(v)
			if desc:
				lines.append('\t// {}'.format(desc))
			lines.append('\t{} {} `json:"{},omitempty"`'.format(go_name, typedef, k))
		lines.append('}')

		return '\n'.join(lines)
	elif '^[0-9]+$' in o.get('patternProperties', {}):
		# See messages.delete
		return 'ArrayOnMeth'

JSON_CACHE = {}

def get_json_cached(file):
	if file == '':
		file = 'objects.json'
	if file not in JSON_CACHE:
		with open('vk-api-schema/' + file) as f:
			JSON_CACHE[file] = json.load(f)

	return JSON_CACHE[file]

def traverse_dict(obj, path):
	for elem in path.split('/'):
		if not elem:
			continue

		obj = obj.get(elem)

		if obj is None:
			break

	return obj

def goify_ref(r):
	file, path = r.split('#')
	last_part = path.split('/')[-1]

	if last_part in KNOWN_TYPE_REFS:
		return 'vk.' + KNOWN_TYPE_REFS[last_part], None

	resolved = resolve_ref(r)
	if resolved is not None:
		return goify_type(resolved)

	return 'genTODOType /* {} */'.format(r), None

def resolve_ref(r):
	file, path = r.split('#')
	last_part = path.split('/')[-1]

	if file == '':
		file = 'objects.json'
	
	if file == 'responses.json' or last_part in FORCE_CODE_GENERATE:
		j = get_json_cached(file)
		return traverse_dict(j, path)
	
def goify_type(resp):
	slice_level = 0
	desc = resp.get('description')

	if 'allOf' in resp:
		lines = ['struct {']
		for x in resp['allOf']:
			t, _ = goify_type(x)
			if t.startswith('struct'):
				t = t.strip().lstrip('struct {').rstrip('}')
			lines.append(t)
		lines.append('}')

		return '\n'.join(lines), None

	while True:
		t = resp.get('type') or resp.get('$ref')

		if t is None and 'properties' in resp:
			t = 'object'

		if t is None:
			print('is this oneOf/allOf?')
			return 'genTODOType', None

		if t != 'array':
			break

		slice_level += 1
		resp = resp['items']

	new_desc = resp.get('description')
	if new_desc:
		desc = new_desc

	if isinstance(t, list):
		if set(t) == {'integer', 'string'}:
			#resp_typedef = 'IntOrString'
			resp_typedef = 'genTODOType'
		else:
			print('type set but not int/string')
			return 'UnexpectedType2', None
	elif t in NATIVE_TYPES:
		resp_typedef = NATIVE_TYPES[t]
	elif t == 'object':
		resp_typedef = goify_object(resp)
	elif '#/definitions' in t:
		resp_typedef, new_desc = goify_ref(t)
		if new_desc:
			desc = new_desc
	else:
		print('Unexpected type', t)
		return 'UnexpectedType3', None

	if not isinstance(resp_typedef, str):
		return 'UnexpectedType4', None

	if slice_level != 0:
		resp_typedef = '[]'*slice_level + resp_typedef

	return resp_typedef, desc

def goify_resp(resp):
	while '$ref' in resp:
		resp = resolve_ref(resp['$ref'])

	assert(resp['type'] == 'object')
	resp = resp['properties']['response']
	return goify_type(resp)

def parse_type(name, decl, varname):
	if decl in NATIVE_TYPES_ZERO_VALS:
		return_type = name
		zero_val, parser_statement = NATIVE_TYPES_ZERO_VALS[decl]
		parser_statement = parser_statement.format(varname, name)
		return_val = varname
		interface_val = '&' + varname
	elif decl.startswith('[]'):
		return_type = name
		zero_val = 'nil'
		return_val = varname
		interface_val = '&' + varname
		parser_statement = 'err = json.Unmarshal(r, {})'.format(interface_val)
	elif decl.startswith('interface'):
		return_type = name
		zero_val = 'nil'
		return_val = varname
		interface_val = varname
		parser_statement = 'err = json.Unmarshal(r, {})'.format(interface_val)
	else:
		return_type = '*' + name
		zero_val = 'nil'
		return_val = '&' + varname
		interface_val = '&' + varname
		parser_statement = 'err = json.Unmarshal(r, {})'.format(interface_val)

	return return_type, zero_val, return_val, interface_val, parser_statement

for namespace, ns_methods in methods.items():
	if namespace in {'ads',}:
		# it's in beta anyway, and has too much shit to be handled
		continue

	go_ns = goify_namespace(namespace)
	f = open('api' + go_ns + '.go', 'w')
	def writeln(s=None):
		if s is not None:
			f.write(str(s))
		f.write('\n')

	writeln('package vkapi\n')

	writeln('import (')
	writeln('\t"strconv"')
	writeln('\t"encoding/json"')
	writeln('\t"github.com/stek29/vk"')
	writeln(')\n')

	writeln('// {} implements VK API namespace `{}`'.format(go_ns, namespace))
	writeln('type {} struct {{'.format(go_ns))
	writeln('\tAPI vk.API')
	writeln('}\n')

	for mtd in ns_methods:
		print('doin', mtd['name'])
		go_mtd_name = goify_method(mtd['name'])

		res = mtd['responses']['response']
		extref = mtd['responses'].get('extendedResponse')

		has_params = len(mtd.get('parameters', [])) != 0

		if has_params:
			writeln('// {}{}Params are params for {}.{}'.format(
				go_ns,
				go_mtd_name,
				go_ns,
				go_mtd_name
			))

			writeln('type {}{}Params struct {{'.format(go_ns, go_mtd_name))
			for param in mtd['parameters']:
				try:
					writeln('// {}'.format(param['description']))
				except KeyError:
					pass
					# print('No description for {}:{}'.format(mtd['name'], param['name']))

				writeln(goify_field(param))
			writeln('}\n')

		if '#/definitions/ok_response' in res.get('$ref', ''):
			return_type = resp_typename = 'bool'
			zero_val = 'false'
			return_val = 'resp'
			interface_val = '&resp'
			resp_parser_trailer = '\treturn decodeBoolIntResponse(r)'
		else:
			res_goified, res_desc = goify_resp(res)
			resp_typename = go_ns + go_mtd_name + 'Response'
			trailer_lst = [
				'\tvar resp {}'.format(resp_typename),
			]

			if extref is not None:
				extres_goified, extres_desc = goify_resp(extref)

				_, _, _, nrm_return_val, nrm_parser_statement = parse_type(resp_typename + 'Normal', res_goified, 'tmp')
				_, _, _, ext_return_val, ext_parser_statement = parse_type(resp_typename + 'Extended', extres_goified, 'tmp')
				nrm_resp_typename = resp_typename + 'Normal'
				ext_resp_typename = resp_typename + 'Extended'

				return_type, zero_val, return_val, interface_val, parser_statement = parse_type(resp_typename, 'interface{}', 'resp')

				trailer_lst += [
					'if params.Extended {',
					'\tvar tmp {}'.format(ext_resp_typename),
					'\t{}'.format(ext_parser_statement),
					'\tresp = {}'.format(ext_return_val),
					'} else {',
					'\tvar tmp {}'.format(nrm_resp_typename),
					'\t{}'.format(nrm_parser_statement),
					'\tresp = {}'.format(nrm_return_val),
					'}'
				]
			else:
				return_type, zero_val, return_val, interface_val, parser_statement = parse_type(resp_typename, res_goified, 'resp')
				trailer_lst.append(
					parser_statement
				)

			resp_parser_trailer = '\n\t'.join(trailer_lst + [
				'if err != nil {',
				'\treturn {}, err'.format(zero_val),
				'}',
				'return {}, nil'.format(return_val),
			])

			writeln('// {}{}Response is response for {}.{}'.format(
				go_ns,
				go_mtd_name,
				go_ns,
				go_mtd_name
			))

			if extref is None:
				if res_desc:
					writeln('// {}'.format(res_desc))
				if res_goified not in NATIVE_TYPES_ZERO_VALS:
					writeln('//easyjson:json')
				writeln('type {}{}Response {}'.format(go_ns, go_mtd_name, res_goified))
			else:
				writeln('// Either {}{}ResponseNormal or {}{}ResponseExtended, depending on Extended flag'.format(
					go_ns, go_mtd_name, go_ns, go_mtd_name
				))
				
				writeln('type {}{}Response interface{{'.format(go_ns, go_mtd_name))
				writeln('\tis{}{}()'.format(go_ns, go_mtd_name))
				writeln('}\n')

				writeln('// {}{}ResponseNormal is non-extended version of {}{}Response'.format(
					go_ns, go_mtd_name, go_ns, go_mtd_name
				))

				if res_desc:
					writeln('// {}'.format(res_desc))

				if res_goified not in NATIVE_TYPES_ZERO_VALS:
					writeln('//easyjson:json')
				writeln('type {}{}ResponseNormal {}'.format(go_ns, go_mtd_name, res_goified))

				writeln('\nfunc ({}{}ResponseNormal) is{}{}(){{}}\n'.format(
					go_ns, go_mtd_name, go_ns, go_mtd_name
				))

				writeln('// {}{}ResponseExtended is extended version of {}{}Response'.format(
					go_ns, go_mtd_name, go_ns, go_mtd_name
				))

				if extres_desc:
					writeln('// {}'.format(extres_desc))
				if extres_goified not in NATIVE_TYPES_ZERO_VALS:
					writeln('//easyjson:json')
				writeln('type {}{}ResponseExtended {}'.format(go_ns, go_mtd_name, extres_goified))

				writeln('\nfunc ({}{}ResponseExtended) is{}{}(){{}}\n'.format(
					go_ns, go_mtd_name, go_ns, go_mtd_name
				))

		if 'description' in mtd:
			mtd_desc = mtd['description']
		else:
			#print('No description for {}'.format(mtd['name']))
			mtd_desc = 'does {}'.format(mtd['name'])

		writeln('// {} {}'.format(go_mtd_name, mtd_desc))
		writeln('func (v {}) {} ({}) ({}, error) {{'.format(
			go_ns,
			go_mtd_name,
			'params {}Params'.format(go_ns + go_mtd_name) if has_params else '',
			return_type
		))

		writeln('\tr, err := v.API.Request("{}", {})'.format(mtd['name'], 'params' if has_params else 'nil'))
		writeln('\tif err != nil {')
		writeln('\t\treturn {}, err'.format(zero_val))
		writeln('\t}\n')

		writeln(resp_parser_trailer)

		writeln('}\n')

	f.close()
	os.system('goimports -w "api{}.go"'.format(go_ns))
